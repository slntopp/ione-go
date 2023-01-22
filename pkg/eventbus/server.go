package eventbus

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/slntopp/nocloud-proto/access"
	pb "github.com/slntopp/nocloud-proto/events"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	RabbitMQConn string
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("RABBITMQ_CONN", "amqp://nocloud:secret@rabbitmq:5672/")
	RabbitMQConn = viper.GetString("RABBITMQ_CONN")
}

type EventBusServer struct {
	pb.UnimplementedEventsServiceServer
	log  *zap.Logger
	bus  *EventBus
	db   driver.Database
	rbmq *amqp.Connection

	ctrl    graph.AccountsController
	ns_ctrl graph.NamespacesController
}

func NewServer(logger *zap.Logger, conn *amqp091.Connection, db driver.Database) *EventBusServer {

	log := logger.Named("EventBusServer")

	log.Info("creating new EvenBusServer instance")

	bus, err := NewEventBus(conn, log)
	if err != nil {
		log.Fatal("cannot create EventBus", zap.Error(err))
	}

	return &EventBusServer{
		log:  log,
		bus:  bus,
		db:   db,
		rbmq: conn,
		ctrl: graph.NewAccountsController(
			log.Named("AccountsController"), db,
		),
		ns_ctrl: graph.NewNamespacesController(
			log.Named("NamespacesController"), db,
		),
	}
}

func (s *EventBusServer) ListenBusQueue(ctx context.Context) {
	log := s.log.Named("Bus queue listener")
init:
	ch, err := s.rbmq.Channel()
	if err != nil {
		log.Error("Failed to open a channel", zap.Error(err))
		time.Sleep(time.Second)
		goto init
	}

	queue, _ := ch.QueueDeclare(
		"events",
		true, false, false, true, nil,
	)

	events, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Error("Failed to register a consumer", zap.Error(err))
		time.Sleep(time.Second)
		goto init
	}

	for msg := range events {
		log.Debug("Received a message")
		var event pb.Event
		err = proto.Unmarshal(msg.Body, &event)
		if err != nil {
			log.Error("Failed to unmarshal event", zap.Error(err))
			if err = msg.Ack(false); err != nil {
				log.Warn("Failed to Acknowledge the delivery while unmarshal message", zap.Error(err))
			}
			continue
		}

		handler, ok := handlers[event.Key]
		if !ok {
			log.Warn("Handler not fount", zap.String("handler", event.Key))
			if err = msg.Ack(false); err != nil {
				log.Warn("Failed to Acknowledge the delivery while unmarshal message", zap.Error(err))
			}
			continue
		}

		updEvent, err := handler(ctx, &event, s.db)
		if err != nil {
			log.Error("Fail to call handler", zap.Any("handler type", event.Key), zap.String("err", err.Error()))
			if err = msg.Ack(false); err != nil {
				log.Warn("Failed to Acknowledge the delivery while unmarshal message", zap.Error(err))
			}
			continue
		}

		_, err = s.Publish(ctx, updEvent)
		if err != nil {
			log.Error("Failed to publish upd event", zap.String("err", err.Error()))
			if err = msg.Ack(false); err != nil {
				log.Warn("Failed to Acknowledge the delivery while unmarshal message", zap.Error(err))
			}
			continue
		}

		if err = msg.Ack(false); err != nil {
			log.Warn("Failed to Acknowledge the delivery while unmarshal message", zap.Error(err))
		}
	}
}

func (s *EventBusServer) Publish(ctx context.Context, event *pb.Event) (*pb.Response, error) {
	log := s.log.Named("Publish")
	log.Debug("Request received", zap.Any("event", event))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	ns, err := s.ns_ctrl.Get(ctx, schema.ROOT_NAMESPACE_KEY)
	if err != nil {
		return nil, err
	}
	ok := graph.HasAccess(ctx, s.db, requestor, ns.ID, access.Level_ADMIN)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights to perform Event publish")
	}

	event.Id = uuid.New().String()
	if err := s.bus.Pub(ctx, event); err != nil {
		return nil, err
	}

	return &pb.Response{}, nil
}

func (s *EventBusServer) Consume(req *pb.ConsumeRequest, srv pb.EventsService_ConsumeServer) error {
	log := s.log.Named("Consume")
	log.Info("Request received", zap.Any("request", req))

	acc, err := s.ctrl.Get(srv.Context(), req.Uuid)
	if err != nil {
		log.Warn("Error getting Account", zap.String("account", req.Uuid), zap.Error(err))
		return status.Error(codes.Internal, "Not enough Access rights")
	}
	if acc.Access == nil || acc.Access.Level < access.Level_ADMIN {
		return status.Error(codes.PermissionDenied, "Not enough Access rights")
	}

	ch, err := s.bus.Sub(srv.Context(), req)
	if err != nil {
		return err
	}
	defer s.bus.Unsub(req)

	done := srv.Context().Done()

	go func() {
		for msg := range ch {
			srv.Send(msg)
		}
	}()

	<-done

	return nil
}

func (s *EventBusServer) List(ctx context.Context, req *pb.ConsumeRequest) (*pb.Events, error) {
	log := s.log.Named("List")
	log.Info("Request received", zap.Any("request", req))

	acc, err := s.ctrl.Get(ctx, req.Uuid)
	if err != nil {
		log.Warn("Error getting Account", zap.String("account", req.Uuid), zap.Error(err))
		return nil, status.Error(codes.Internal, "Not enough Access rights")
	}
	if acc.Access == nil || acc.Access.Level < access.Level_ADMIN {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access rights")
	}

	events, err := s.bus.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.Events{Events: events}, nil
}

func (s *EventBusServer) Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.Response, error) {
	log := s.log.Named("Cancel")
	log.Info("Request received", zap.Any("request", req))

	acc, err := s.ctrl.Get(ctx, req.Uuid)
	if err != nil {
		log.Warn("Error getting Account", zap.String("account", req.Uuid), zap.Error(err))
		return nil, status.Error(codes.Internal, "Not enough Access rights")
	}
	if acc.Access == nil || acc.Access.Level < access.Level_ADMIN {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access rights")
	}

	return &pb.Response{}, s.bus.Cancel(ctx, req)
}
