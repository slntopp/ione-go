package eventbus

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	pb "github.com/slntopp/nocloud-proto/events"
	"go.uber.org/zap"
)

type EventBusServer struct {
	pb.UnimplementedEventsServiceServer
	log *zap.Logger
	bus *EventBus
}

func NewServer(logger *zap.Logger, conn *amqp091.Connection) *EventBusServer {

	log := logger.Named("EventBusServer")

	log.Info("creating new EvenBusServer instance")

	bus, err := NewEventBus(conn, log)
	if err != nil {
		log.Fatal("cannot create EventBus", zap.Error(err))
	}

	return &EventBusServer{
		log: log,
		bus: bus,
	}
}

func (s *EventBusServer) Publish(ctx context.Context, event *pb.Event) (*pb.Response, error) {

	s.log.Info("got publish request")

	if err := s.bus.Pub(ctx, event); err != nil {
		return nil, err
	}

	return &pb.Response{}, nil
}

func (s *EventBusServer) Consume(req *pb.ConsumeRequest, srv pb.EventsService_ConsumeServer) error {

	s.log.Info("got consume request")

	ch, err := s.bus.Sub(req.Key)
	if err != nil {
		return err
	}

	for msg := range ch {
		srv.Send(msg)
	}

	return nil
}
