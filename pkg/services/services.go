/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

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
package services

import (
	"context"
	"fmt"

	"github.com/arangodb/go-driver"
	driverpb "github.com/slntopp/nocloud/pkg/drivers/instance/vanilla"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/instances/proto"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/access"
	"github.com/slntopp/nocloud/pkg/nocloud/roles"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	pb "github.com/slntopp/nocloud/pkg/services/proto"
	stpb "github.com/slntopp/nocloud/pkg/states/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Routine struct {
	Name string
	LastExec string
	Running bool
}

type ServicesServer struct {
	pb.UnimplementedServicesServiceServer
	db      driver.Database
	ctrl    graph.ServicesController
	sp_ctrl graph.ServicesProvidersController
	ns_ctrl graph.NamespacesController

	drivers  map[string]driverpb.DriverServiceClient
	states stpb.StatesServiceClient

	monitoring Routine

	log *zap.Logger
}

func NewServicesServer(log *zap.Logger, db driver.Database, gc stpb.StatesServiceClient) *ServicesServer {
	return &ServicesServer{
		log: log, db: db, ctrl: graph.NewServicesController(log, db),
		sp_ctrl:  graph.NewServicesProvidersController(log, db),
		ns_ctrl:  graph.NewNamespacesController(log, db),
		drivers:  make(map[string]driverpb.DriverServiceClient),
		states: gc,
		monitoring: Routine{
			Name: "Monitoring",
			Running: false,
		},
	}
}

type InstancesGroupDriverContext struct {
	sp     *graph.ServicesProvider
	client *driverpb.DriverServiceClient
}

func (s *ServicesServer) RegisterDriver(type_key string, client driverpb.DriverServiceClient) {
	s.drivers[type_key] = client
}

func (s *ServicesServer) DoTestServiceConfig(ctx context.Context, log *zap.Logger, request *pb.CreateRequest) (*pb.TestConfigResponse, *graph.Namespace, error) {
	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	response := &pb.TestConfigResponse{Result: true, Errors: make([]*pb.TestConfigError, 0)}

	namespace, err := s.ns_ctrl.Get(ctx, request.GetNamespace())
	if err != nil {
		s.log.Debug("Error getting namespace", zap.Error(err))
		return nil, nil, status.Error(codes.NotFound, "Namespace not found")
	}
	// Checking if requestor has access to Namespace Service going to be put in
	ok := graph.HasAccess(ctx, s.db, requestor, namespace.ID.String(), access.ADMIN)
	if !ok {
		return nil, nil, status.Error(codes.PermissionDenied, "Not enough access rights to Namespace")
	}

	service := request.GetService()
	groups := service.GetInstancesGroups()

	log.Debug("Init validation", zap.Any("groups", groups), zap.Int("amount", len(groups)))
	for _, group := range service.GetInstancesGroups() {
		log.Debug("Validating Instances Group", zap.String("group", group.Title))
		groupType := group.GetType()

		config_err := pb.TestConfigError{
			InstanceGroup: group.Title,
		}

		client, ok := s.drivers[groupType]
		if !ok {
			response.Result = false
			config_err.Error = fmt.Sprintf("Driver Type '%s' not registered", groupType)
			response.Errors = append(
				response.Errors, &config_err,
			)
			continue
		}

		res, err := client.TestInstancesGroupConfig(ctx, &proto.TestInstancesGroupConfigRequest{Group: group})
		if err != nil {
			response.Result = false
			config_err.Error = fmt.Sprintf("Error validating group '%s'", group.Title)
			response.Errors = append(
				response.Errors, &config_err,
			)
			continue
		}
		if !res.GetResult() {
			response.Result = false
			errors := make([]*pb.TestConfigError, 0)
			for _, confErr := range res.Errors {
				errors = append(errors, &pb.TestConfigError{
					Error:         confErr.Error,
					Instance:      confErr.Instance,
					InstanceGroup: group.Title,
				})
			}
			response.Errors = append(response.Errors, errors...)
			continue
		}
		log.Debug("Validated Instances Group", zap.String("group", group.Title))
	}

	return response, &namespace, nil
}

func (s *ServicesServer) TestConfig(ctx context.Context, request *pb.CreateRequest) (*pb.TestConfigResponse, error) {
	log := s.log.Named("TestServiceConfig")
	log.Debug("Request received", zap.Any("request", request))
	response, _, err := s.DoTestServiceConfig(ctx, log, request)
	return response, err
}

func (s *ServicesServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.Service, error) {
	log := s.log.Named("CreateService")
	log.Debug("Request received", zap.Any("request", request))
	testResult, namespace, err := s.DoTestServiceConfig(ctx, log, request)

	if err != nil {
		return nil, err
	} else if !testResult.Result {
		return nil, status.Error(codes.InvalidArgument, "Config didn't pass test")
	}

	service := request.GetService()
	doc, err := s.ctrl.Create(ctx, service)
	if err != nil {
		log.Error("Error while creating service", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating Service")
	}

	err = s.ctrl.Join(ctx, doc, namespace, access.ADMIN, roles.OWNER)
	if err != nil {
		log.Error("Error while joining service to namespace", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while joining service to namespace")
	}
	return service, nil
}

func (s *ServicesServer) Up(ctx context.Context, request *pb.UpRequest) (*pb.UpResponse, error) {
	log := s.log.Named("Up")
	log.Debug("Request received", zap.Any("request", request))

	service, err := s.ctrl.Get(ctx, request.GetUuid())
	if err != nil {
		log.Debug("Error getting Service", zap.Error(err))
		return nil, status.Error(codes.NotFound, "Service not found")
	}
	log.Debug("Found Service", zap.Any("service", service))

	deploy_policies := request.GetDeployPolicies()
	contexts := make(map[string]*InstancesGroupDriverContext)

	for _, group := range service.GetInstancesGroups() {
		sp_id := deploy_policies[group.GetUuid()]
		sp, err := s.sp_ctrl.Get(ctx, sp_id)
		if err != nil {
			log.Error("Error getting ServiceProvider", zap.Error(err), zap.String("id", sp_id))
			return nil, status.Errorf(codes.InvalidArgument, "Error getting ServiceProvider(%s)", sp_id)
		}

		groupType := group.GetType()
		client, ok := s.drivers[groupType]
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Driver of type '%s' not registered", groupType)
		}
		contexts[group.GetUuid()] = &InstancesGroupDriverContext{sp, &client}
	}

	err = s.ctrl.SetStatus(ctx, service, pb.ServiceStatus_PROC)
	if err != nil {
		log.Error("Error updating Service", zap.Error(err), zap.Any("service", service))
		return nil, status.Error(codes.Internal, "Error storing updates")
	}

	result := &pb.UpResponse{Errors: make([]*pb.UpError, 0)}
	for _, group := range service.GetInstancesGroups() {
		c, ok := contexts[group.GetUuid()]
		if !ok {
			log.Debug("Instance Group has no context", zap.String("group", group.GetUuid()), zap.String("service", service.GetUuid()))
			continue
		}
		client := *c.client
		sp := c.sp

		response, err := client.Up(ctx, &driverpb.UpRequest{Group: group, ServicesProvider: sp.ServicesProvider})
		if err != nil {
			log.Error("Error deploying group", zap.Any("service_provider", sp), zap.Any("group", group), zap.Error(err))
			result.Errors = append(result.Errors, &pb.UpError{
				Data: map[string]string{
					"group": group.GetUuid(),
					"error": err.Error(),
				},
			})
			continue
		}
		log.Debug("Up Request Result", zap.Any("response", response))

		// TODO: Change to Hash comparation
		// TODO: Add cleanups
		if len(group.Instances) != len(response.GetGroup().GetInstances()) {
			log.Error("Instances config changed by Driver")
			result.Errors = append(result.Errors, &pb.UpError{
				Data: map[string]string{
					"group": group.GetUuid(),
					"error": "Instances config changed by Driver",
				},
			})
			continue
		}
		for i, instance := range response.GetGroup().GetInstances() {
			group.Instances[i].Data = instance.GetData()
		}

		group.Data = response.GetGroup().GetData()
		err = s.ctrl.IGController().Provide(ctx, group.Uuid, sp.Uuid)
		if err != nil {
			log.Error("Error linking group to ServiceProvider", zap.Any("service_provider", sp.GetUuid()), zap.Any("group", group), zap.Error(err))
			result.Errors = append(result.Errors, &pb.UpError{
				Data: map[string]string{
					"group": group.GetUuid(),
					"error": err.Error(),
				},
			})
			continue
		}
		log.Debug("Updated Group", zap.Any("group", group))
	}

	service.Status = pb.ServiceStatus_UP
	log.Debug("Updated Service", zap.Any("service", service))
	err = s.ctrl.Update(ctx, service, false)
	if err != nil {
		log.Error("Error updating Service", zap.Error(err), zap.Any("service", service))
		return nil, status.Error(codes.Internal, "Error storing updates")
	}

	return &pb.UpResponse{}, nil
}

func (s *ServicesServer) Down(ctx context.Context, request *pb.DownRequest) (*pb.DownResponse, error) {
	log := s.log.Named("Down")
	log.Debug("Request received", zap.Any("request", request))

	service, err := s.ctrl.Get(ctx, request.GetUuid())
	if err != nil {
		log.Debug("Error getting Service", zap.Error(err))
		return nil, status.Error(codes.NotFound, "Service not found")
	}
	log.Debug("Found Service", zap.Any("service", service))

	contexts := make(map[string]*InstancesGroupDriverContext)

	for _, group := range service.GetInstancesGroups() {
		if group.Sp == nil {
			log.Debug("Group is unprovisioned, skipping", zap.String("group", group.GetUuid()))
			continue
		}

		sp, err := s.sp_ctrl.Get(ctx, *group.Sp)
		if err != nil {
			log.Error("Error getting ServiceProvider", zap.Error(err), zap.String("id", *group.Sp))
			return nil, status.Errorf(codes.InvalidArgument, "Error getting ServiceProvider(%s)", *group.Sp)
		}

		groupType := group.GetType()
		client, ok := s.drivers[groupType]
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Driver of type '%s' not registered", groupType)
		}

		contexts[group.GetUuid()] = &InstancesGroupDriverContext{sp, &client}
	}

	err = s.ctrl.SetStatus(ctx, service, pb.ServiceStatus_PROC)
	if err != nil {
		log.Error("Error updating Service", zap.Error(err), zap.Any("service", service))
		return nil, status.Error(codes.Internal, "Error storing updates")
	}

	for i, group := range service.GetInstancesGroups() {
		c, ok := contexts[group.GetUuid()]
		if !ok {
			log.Debug("Instance Group has no context, i.e. provision", zap.String("group", group.GetUuid()), zap.String("service", service.GetUuid()))
			continue
		}
		client := *c.client
		sp := c.sp

		res, err := client.Down(ctx, &driverpb.DownRequest{Group: group, ServicesProvider: sp.ServicesProvider})
		if err != nil {
			log.Error("Error undeploying group", zap.Any("service_provider", sp), zap.Any("group", group), zap.Error(err))
			continue
		}
		group := res.GetGroup()
		// err = s.ctrl.Unprovide(ctx, group.GetUuid())
		// if err != nil {
		// 	log.Error("Error unlinking group from ServiceProvider", zap.Any("service_provider", sp.GetUuid()), zap.Any("group", group), zap.Error(err))
		// 	continue
		// }
		service.InstancesGroups[i] = group
	}

	err = s.ctrl.SetStatus(ctx, service, pb.ServiceStatus_INIT)
	if err != nil {
		log.Error("Error updating Service", zap.Error(err), zap.Any("service", service))
		return nil, status.Error(codes.Internal, "Error storing updates")
	}

	return &pb.DownResponse{}, nil
}

func (s *ServicesServer) Get(ctx context.Context, request *pb.GetRequest) (res *pb.Service, err error) {
	log := s.log.Named("Get")
	log.Debug("Request received", zap.Any("request", request))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	r, err := s.ctrl.Get(ctx, request.GetUuid())
	if err != nil {
		log.Debug("Error getting Service from DB", zap.Error(err))
		return nil, status.Error(codes.NotFound, "Service not Found in DB")
	}

	ok := graph.HasAccess(ctx, s.db, requestor, driver.NewDocumentID(schema.SERVICES_COL, r.Uuid).String(), access.READ)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights")
	}

	states, err := s.GetStatesInternal(ctx, r)
	if err != nil {
		log.Error("Error getting Instances States", zap.String("uuid", r.GetUuid()), zap.Error(err))
		return r, nil
	}
	log.Debug("Got Instances States", zap.Any("states", states))

	for _, group := range r.GetInstancesGroups() {
		for _, inst := range group.GetInstances() {
			inst.State = states.States[inst.GetUuid()]
		}
	}

	return r, nil
}

func (s *ServicesServer) List(ctx context.Context, request *pb.ListRequest) (response *pb.Services, err error) {
	log := s.log.Named("List")
	log.Debug("Request received", zap.String("namespace", request.GetNamespace()), zap.String("show_deleted", request.GetShowDeleted()))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	r, err := s.ctrl.List(ctx, requestor, request)
	if err != nil {
		log.Debug("Error reading Services from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error reading Services from DB")
	}

	return &pb.Services{Pool: r}, nil
}

func (s *ServicesServer) Delete(ctx context.Context, request *pb.DeleteRequest) (response *pb.DeleteResponse, err error) {
	log := s.log.Named("Delete")
	log.Debug("Request received", zap.Any("request", request))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	r, err := s.ctrl.Get(ctx, request.GetUuid())
	if err != nil {
		log.Debug("Error getting Service from DB", zap.Error(err))
		return nil, status.Error(codes.NotFound, "Service not Found in DB")
	}

	ok := graph.HasAccess(ctx, s.db, requestor, driver.NewDocumentID(schema.SERVICES_COL, r.Uuid).String(), access.MGMT)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights")
	}

	err = s.ctrl.Delete(ctx, r)
	if err != nil {
		log.Error("Error Deleting Service", zap.Error(err))
		return &pb.DeleteResponse{Result: false, Error: err.Error()}, nil
	}

	return &pb.DeleteResponse{Result: true}, nil
}