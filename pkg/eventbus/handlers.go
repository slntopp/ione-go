package eventbus

import (
	"context"
	"github.com/arangodb/go-driver"
	pb "github.com/slntopp/nocloud-proto/events"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"google.golang.org/protobuf/types/known/structpb"
)

type EventHandler func(context.Context, *pb.Event, driver.Database) (*pb.Event, error)

var handlers = map[string]EventHandler{
	"instance_suspended":          GetInstAccountHandler,
	"instance_unsuspended":        GetInstAccountHandler,
	"instance_created":            GetInstAccountHandler,
	"instance_deleted":            GetInstAccountHandler,
	"expiry_notification":         GetInstAccountHandler,
	"suspend_expiry_notification": GetInstAccountHandler,
	"suspend_delete_instance":     GetInstAccountHandler,
}

var getInstanceAccount = `
LET doc = DOCUMENT(@inst)

LET srv = LAST(
FOR node, edge, path IN 2
    INBOUND doc
    GRAPH @permissions
    FILTER path.edges[*].role == ["owner","owner"]
    FILTER IS_SAME_COLLECTION(node, @@services)
        RETURN node
    )

LET account = LAST(
    FOR node, edge, path IN 2
    INBOUND srv
    GRAPH @permissions
    FILTER path.edges[*].role == ["owner","owner"]
    FILTER IS_SAME_COLLECTION(node, @@accounts)
        RETURN node
    )
    
RETURN {
	account: account._key, 
	service: srv.title, 
	instance: doc.title, 
	product: doc.product, 
	next_payment_date: doc.data.next_payment_date,
	ip: doc.state.meta.networking.public[0]
}
`

type EventInfo struct {
	Account         string  `json:"account"`
	Service         string  `json:"service"`
	Instance        string  `json:"instance"`
	Product         string  `json:"product,omitempty"`
	Ip              string  `json:"ip,omitempty"`
	NextPaymentDate float64 `json:"next_payment_date,omitempty"`
}

func GetInstAccountHandler(ctx context.Context, event *pb.Event, db driver.Database) (*pb.Event, error) {
	if event.GetData() == nil {
		event.Data = make(map[string]*structpb.Value)
	}

	inst := driver.NewDocumentID(schema.INSTANCES_COL, event.GetUuid())

	cursor, err := db.Query(ctx, getInstanceAccount, map[string]interface{}{
		"inst":        inst,
		"permissions": schema.PERMISSIONS_GRAPH.Name,
		"@services":   schema.SERVICES_COL,
		"@accounts":   schema.ACCOUNTS_COL,
	})
	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	var eventInfo EventInfo
	for cursor.HasMore() {
		_, err := cursor.ReadDocument(ctx, &eventInfo)
		if err != nil {
			return nil, err
		}
	}

	event.Data["service"] = structpb.NewStringValue(eventInfo.Service)
	event.Data["instance"] = structpb.NewStringValue(eventInfo.Instance)
	if eventInfo.Product != "" {
		event.Data["product"] = structpb.NewStringValue(eventInfo.Product)
	}
	if eventInfo.Ip != "" {
		event.Data["ip"] = structpb.NewStringValue(eventInfo.Ip)
	}
	if eventInfo.NextPaymentDate != 0 {
		event.Data["next_payment_date"] = structpb.NewNumberValue(eventInfo.NextPaymentDate)
	}
	event.Data["instance_uuid"] = structpb.NewStringValue(event.GetUuid())
	event.Uuid = eventInfo.Account
	event.Type = "email"

	return event, nil
}
