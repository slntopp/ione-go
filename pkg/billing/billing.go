/*
Copyright © 2021-2023 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package billing

import (
	"context"

	"github.com/slntopp/nocloud/pkg/nocloud/schema"

	"github.com/arangodb/go-driver"
	"github.com/slntopp/nocloud-proto/access"
	pb "github.com/slntopp/nocloud-proto/billing"
	healthpb "github.com/slntopp/nocloud-proto/health"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Routine struct {
	Name     string
	LastExec string
	Running  bool
}

type BillingServiceServer struct {
	pb.UnimplementedBillingServiceServer

	log *zap.Logger

	nss          graph.NamespacesController
	plans        graph.BillingPlansController
	transactions graph.TransactionsController
	records      graph.RecordsController
	currencies   graph.CurrencyController
	accounts     graph.AccountsController

	db driver.Database

	gen  *healthpb.RoutineStatus
	proc *healthpb.RoutineStatus
	sus  *healthpb.RoutineStatus
}

func NewBillingServiceServer(logger *zap.Logger, db driver.Database) *BillingServiceServer {
	log := logger.Named("BillingService")
	return &BillingServiceServer{
		log:          log,
		nss:          graph.NewNamespacesController(log, db),
		plans:        graph.NewBillingPlansController(log.Named("PlansController"), db),
		transactions: graph.NewTransactionsController(log.Named("TransactionsController"), db),
		records:      graph.NewRecordsController(log.Named("RecordsController"), db),
		currencies:   graph.NewCurrencyController(log.Named("CurrenciesController"), db),
		accounts:     graph.NewAccountsController(log.Named("AccountsController"), db),
		db:           db,
		gen: &healthpb.RoutineStatus{
			Routine: "Generate Transactions",
			Status: &healthpb.ServingStatus{
				Service: "Billing Machine",
				Status:  healthpb.Status_STOPPED,
			},
		}, proc: &healthpb.RoutineStatus{
			Routine: "Process Transactions",
			Status: &healthpb.ServingStatus{
				Service: "Billing Machine",
				Status:  healthpb.Status_STOPPED,
			},
		},
		sus: &healthpb.RoutineStatus{
			Routine: "Suspend Monitoring",
			Status: &healthpb.ServingStatus{
				Service: "Billing Machine",
				Status:  healthpb.Status_STOPPED,
			},
		},
	}
}

func (s *BillingServiceServer) CreatePlan(ctx context.Context, plan *pb.Plan) (*pb.Plan, error) {
	log := s.log.Named("CreatePlan")
	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("request", zap.Any("plan", plan), zap.String("requestor", requestor))

	ns, err := s.nss.Get(ctx, schema.ROOT_NAMESPACE_KEY)
	if err != nil {
		return nil, err
	}
	ok := graph.HasAccess(ctx, s.db, requestor, ns.ID, access.Level_ADMIN)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access rights to manage BillingPlans")
	}

	res, err := s.plans.Create(ctx, plan)
	if err != nil {
		log.Error("Error creating plan", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error creating plan")
	}

	return res.Plan, nil
}

func (s *BillingServiceServer) UpdatePlan(ctx context.Context, plan *pb.Plan) (*pb.Plan, error) {
	log := s.log.Named("UpdatePlan")
	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("request", zap.Any("plan", plan), zap.String("requestor", requestor))

	ns, err := s.nss.Get(ctx, schema.ROOT_NAMESPACE_KEY)
	if err != nil {
		return nil, err
	}
	ok := graph.HasAccess(ctx, s.db, requestor, ns.ID, access.Level_ADMIN)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access rights to manage BillingPlans")
	}

	res, err := s.plans.Update(ctx, plan)
	if err != nil {
		log.Error("Error updating plan", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error updating plan")
	}

	return res.Plan, nil
}

func (s *BillingServiceServer) DeletePlan(ctx context.Context, plan *pb.Plan) (*pb.Plan, error) {
	log := s.log.Named("DeletePlan")
	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("request", zap.Any("plan", plan), zap.String("requestor", requestor))

	ns, err := s.nss.Get(ctx, schema.ROOT_NAMESPACE_KEY)
	if err != nil {
		return nil, err
	}
	ok := graph.HasAccess(ctx, s.db, requestor, ns.ID, access.Level_ADMIN)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access rights to manage BillingPlans")
	}

	planId := driver.NewDocumentID(schema.BILLING_PLANS_COL, plan.GetUuid())

	cursor, err := s.db.Query(ctx, getInstances, map[string]interface{}{
		"permissions":       schema.PERMISSIONS_GRAPH.Name,
		"plan":              planId,
		"@instances_groups": schema.INSTANCES_GROUPS_COL,
		"@instances":        schema.INSTANCES_COL,
	})
	if err != nil {
		log.Error("Error getting instances", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error getting instances")
	}

	if cursor.HasMore() {
		return nil, status.Error(codes.DataLoss, "Сan't delete plan due to related instances")
	}

	err = s.plans.Delete(ctx, plan)
	if err != nil {
		log.Error("Error deleting plan", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error deleting plan")
	}

	return plan, nil
}

func (s *BillingServiceServer) GetPlan(ctx context.Context, plan *pb.Plan) (*pb.Plan, error) {
	log := s.log.Named("GetPlan")
	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("request", zap.Any("plan", plan), zap.String("requestor", requestor))

	p, err := s.plans.Get(ctx, plan)
	if err != nil {
		log.Error("Error getting plan", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error getting plan")
	}

	if p.Public {
		return p.Plan, nil
	}

	ok := graph.HasAccess(ctx, s.db, requestor, p.ID, access.Level_READ)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access rights to manage BillingPlans")
	}

	return p.Plan, nil
}

var getDefaultCurrencyQuery = `
LET cur = LAST(
    FOR i IN Currencies2Currencies
    FILTER (i.to == 0 || i.from == 0) && i.rate == 1
        RETURN i
)

RETURN cur.to == 0 ? cur.from : cur.to
`

func (s *BillingServiceServer) ListPlans(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	log := s.log.Named("ListPlans")

	var requestor string
	if !req.Anonymously {
		requestor = ctx.Value(nocloud.NoCloudAccount).(string)
	}
	log.Debug("Requestor", zap.String("id", requestor))

	plans, err := s.plans.List(ctx, req.SpUuid)
	if err != nil {
		log.Error("Error listing plans", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error listing plans")
	}

	result := make([]*pb.Plan, 0)
	for _, plan := range plans {
		if plan.Public {
			result = append(result, plan.Plan)
			continue
		}
		if req.Anonymously {
			continue
		}
		ok := graph.HasAccess(ctx, s.db, requestor, plan.ID, access.Level_READ)
		if !ok {
			continue
		}
		result = append(result, plan.Plan)
	}

	if !req.Anonymously {
		acc, err := s.accounts.Get(ctx, requestor)
		if err != nil {
			log.Error("Error getting account", zap.Error(err))
			return nil, status.Error(codes.Internal, "Error getting account")
		}

		cur := acc.Account.GetCurrency()

		defaultCur := pb.Currency_NCU

		queryContext := driver.WithQueryCount(ctx)
		res, err := s.db.Query(queryContext, getDefaultCurrencyQuery, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if res.Count() != 0 {
			_, err = res.ReadDocument(ctx, &defaultCur)
			if err != nil {
				log.Error("Failed to get default cur", zap.Error(err))
				return nil, status.Error(codes.Internal, "Failed to get default cur")
			}
		}

		var rate float64

		if cur == defaultCur {
			rate = 1
		} else {
			rate, err = s.currencies.GetExchangeRateDirect(ctx, defaultCur, cur)
			if err != nil {
				log.Error("Error getting rate", zap.Error(err))
				return nil, status.Error(codes.Internal, "Error getting rate")
			}
		}

		for planIndex := range result {
			plan := result[planIndex]

			products := plan.GetProducts()
			for key := range products {
				products[key].Price *= rate
			}
			plan.Products = products

			resources := plan.GetResources()
			for index := range resources {
				resources[index].Price *= rate
			}
			plan.Resources = resources

			result[planIndex] = plan
		}
	}

	return &pb.ListResponse{Pool: result}, nil
}

const getInstances = `
LET igs = (
    FOR node IN 2 INBOUND @plan GRAPH @permissions
    FILTER IS_SAME_COLLECTION(node, @@instances_groups)
    RETURN node._id
)


LET instances = (
	FOR ig in igs
    	FOR node, edge IN 1 OUTBOUND ig GRAPH @permissions
    	FILTER IS_SAME_COLLECTION(node, @@instances)
    	FILTER edge.role == "owner"
    	RETURN node
)

LET plan = DOCUMENT(@plan)

FOR inst in instances
	FILTER inst.billing_plan.uuid == plan._key
	RETURN inst

`
