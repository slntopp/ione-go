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
package registry

import (
	"context"

	"github.com/arangodb/go-driver"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	"github.com/slntopp/nocloud/pkg/credentials"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/access"
	"github.com/slntopp/nocloud/pkg/nocloud/roles"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	servicespb "github.com/slntopp/nocloud/pkg/services/proto"

	accesspb "github.com/slntopp/nocloud/pkg/access"
	pb "github.com/slntopp/nocloud/pkg/registry/proto"
	accountspb "github.com/slntopp/nocloud/pkg/registry/proto/accounts"
	sc "github.com/slntopp/nocloud/pkg/settings/client"
	settingspb "github.com/slntopp/nocloud/pkg/settings/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	servicesClient servicespb.ServicesServiceClient
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("SERVICES_HOST", "services-registry:8000")
	servicesHost := viper.GetString("SERVICES_HOST")

	servicesConn, err := grpc.Dial(servicesHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	servicesClient = servicespb.NewServicesServiceClient(servicesConn)
}

type AccountsServiceServer struct {
	pb.UnimplementedAccountsServiceServer
	db      driver.Database
	ctrl    graph.AccountsController
	ns_ctrl graph.NamespacesController

	log         *zap.Logger
	SIGNING_KEY []byte
}

func NewAccountsServer(log *zap.Logger, db driver.Database) *AccountsServiceServer {
	return &AccountsServiceServer{
		log: log, db: db,
		ctrl: graph.NewAccountsController(
			log.Named("AccountsController"), db,
		),
		ns_ctrl: graph.NewNamespacesController(
			log.Named("NamespacesController"), db,
		),
	}
}

func (s *AccountsServiceServer) SetupSettingsClient(settingsClient settingspb.SettingsServiceClient, internal_token string) {
	sc.Setup(
		s.log, metadata.AppendToOutgoingContext(
			context.Background(), "authorization", "bearer "+internal_token,
		), &settingsClient,
	)
	var settings AccountPostCreateSettings
	if scErr := sc.Fetch(accountPostCreateSettingsKey, &settings, defaultSettings); scErr != nil {
		s.log.Warn("Cannot fetch settings", zap.Error(scErr))
	}
}

const getOwnServices = `
FOR node, edge, path IN 2 OUTBOUND @account
    GRAPH Permissions
    FILTER path.edges[*].role == ["owner","owner"]
    FILTER IS_SAME_COLLECTION(node, Services)
    RETURN node._key
`

func (s *AccountsServiceServer) Suspend(ctx context.Context, req *accountspb.SuspendRequest) (*accountspb.SuspendResponse, error) {
	log := s.log.Named("SuspendAccount")
	log.Debug("Suspend request received", zap.Any("request", req), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	cursor, err := s.db.Query(ctx, getOwnServices, map[string]interface{}{
		"account": req.Uuid,
	})
	if err != nil {
		log.Error("Error Quering Services to Suspend", zap.Error(err))
		return &accountspb.SuspendResponse{Result: false}, err
	}
	defer cursor.Close()

	for {
		var serviceUuid string
		meta, err := cursor.ReadDocument(ctx, &serviceUuid)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Error("Error Reading Service UUID", zap.Error(err), zap.Any("meta", meta))
			continue
		}
		if _, err := servicesClient.Suspend(ctx, &servicespb.SuspendRequest{Uuid: serviceUuid}); err != nil {
			log.Error("Error Suspending Service", zap.Error(err))
		}
	}

	acc, err := s.ctrl.Get(ctx, req.GetUuid())
	if err != nil {
		log.Debug("Error getting account", zap.String("requested_id", req.Uuid), zap.Any("error", err))
		return nil, status.Error(codes.NotFound, "Account not found")
	}

	if err := s.ctrl.Update(ctx, acc, map[string]interface{}{"suspended": true}); err != nil {
		log.Debug("Error updating account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating account")
	}

	log.Debug("Suspend request end")
	return &accountspb.SuspendResponse{Result: true}, nil
}

func (s *AccountsServiceServer) Unsuspend(ctx context.Context, req *accountspb.UnsuspendRequest) (*accountspb.UnsuspendResponse, error) {
	log := s.log.Named("UnsuspendAccount")
	log.Debug("Unsuspend request received", zap.Any("request", req), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	cursor, err := s.db.Query(ctx, getOwnServices, map[string]interface{}{
		"account": req.Uuid,
	})
	if err != nil {
		log.Error("Error Quering Services to Unsuspend", zap.Error(err))
		return &accountspb.UnsuspendResponse{Result: false}, err
	}
	defer cursor.Close()

	for {
		var serviceUuid string
		meta, err := cursor.ReadDocument(ctx, &serviceUuid)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Error("Error Reading Service UUID", zap.Error(err), zap.Any("meta", meta))
			continue
		}
		if _, err := servicesClient.Suspend(ctx, &servicespb.SuspendRequest{Uuid: serviceUuid}); err != nil {
			log.Error("Error Unsuspending Service", zap.Error(err))
		}
	}

	acc, err := s.ctrl.Get(ctx, req.GetUuid())
	if err != nil {
		log.Debug("Error getting account", zap.String("requested_id", req.GetUuid()), zap.Any("error", err))
		return nil, status.Error(codes.NotFound, "Account not found")
	}

	if err := s.ctrl.Update(ctx, acc, map[string]interface{}{"suspended": false}); err != nil {
		log.Debug("Error updating account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating account")
	}

	log.Debug("Unsuspend request end")
	return &accountspb.UnsuspendResponse{Result: true}, nil
}

func (s *AccountsServiceServer) Get(ctx context.Context, request *accountspb.GetRequest) (*accountspb.Account, error) {
	log := s.log.Named("GetAccount")
	log.Debug("Get request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	requested := request.GetUuid()
	if requested == "me" {
		requested = requestor
	}

	log.Debug("Retrieving account", zap.String("uuid", requested))
	acc, err := graph.GetWithAccess[graph.Account](ctx, s.db, driver.NewDocumentID(schema.ACCOUNTS_COL, request.GetUuid()))
	if err != nil || acc.Access == nil {
		log.Debug("Error getting account", zap.Any("error", err))
		return nil, status.Error(codes.NotFound, "Account not found")
	}
	log.Debug("Retrieved account", zap.Any("account", acc))

	// Provide public information without access check
	if request.GetPublic() {
		return &accountspb.Account{Title: acc.Account.GetTitle()}, nil
	}

	if acc.Key == requestor {
		return acc.Account, nil
	}

	if acc.Access.Level < accesspb.Level_READ {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights to Account")
	}

	if acc.Access.Level < accesspb.Level_ROOT {
		acc.SuspendConf = nil
	}

	return acc.Account, nil
}

func (s *AccountsServiceServer) List(ctx context.Context, request *accountspb.ListRequest) (*accountspb.ListResponse, error) {
	log := s.log.Named("ListAccounts")
	log.Debug("List request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	acc, err := s.ctrl.Get(ctx, requestor)
	if err != nil {
		log.Debug("Error getting requestor account", zap.Any("error", err))
		return nil, status.Error(codes.PermissionDenied, "Requestor Account not found")
	}

	pool, err := graph.ListWithAccess[graph.Account](ctx, log, s.db, acc.ID, schema.ACCOUNTS_COL, request.GetDepth())
	if err != nil {
		log.Debug("Error listing accounts", zap.Any("error", err))
		return nil, status.Error(codes.Internal, "Error listing accounts")
	}
	log.Debug("List result", zap.Any("pool", pool))

	result := make([]*accountspb.Account, len(pool))
	for i, acc := range pool {
		if acc.Access.Level < accesspb.Level_ROOT {
			acc.Account.SuspendConf = nil
		}
		result[i] = acc.Account
	}
	log.Debug("Convert result", zap.Any("pool", result))

	return &accountspb.ListResponse{Pool: result}, nil
}

func (s *AccountsServiceServer) Token(ctx context.Context, request *accountspb.TokenRequest) (*accountspb.TokenResponse, error) {
	log := s.log.Named("Token")

	log.Debug("Token request received", zap.Any("request", request))
	acc, ok := s.ctrl.Authorize(ctx, request.Auth.Type, request.Auth.Data...)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Wrong credentials given")
	}
	log.Debug("Authorized user", zap.String("ID", acc.ID.String()))

	claims := jwt.MapClaims{}
	claims[nocloud.NOCLOUD_ACCOUNT_CLAIM] = acc.Key
	claims["exp"] = request.Exp

	if request.GetRootClaim() {
		ns := driver.NewDocumentID(schema.NAMESPACES_COL, "0")
		ok, lvl := graph.AccessLevel(ctx, s.db, acc.Key, ns)
		if !ok {
			lvl = 0
		}
		log.Debug("Adding Root claim to the token", zap.Int32("access_lvl", lvl))
		claims[nocloud.NOCLOUD_ROOT_CLAIM] = lvl
	}

	if sp := request.GetSpClaim(); sp != "" {
		ok, lvl := graph.AccessLevel(ctx, s.db, acc.Key, driver.NewDocumentID(schema.SERVICES_PROVIDERS_COL, sp))
		if !ok {
			lvl = 0
		}
		log.Debug("Adding ServicesProvider claim to the token", zap.Int32("access_lvl", lvl))
		claims[nocloud.NOCLOUD_SP_CLAIM] = lvl
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(s.SIGNING_KEY)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to issue token")
	}

	return &accountspb.TokenResponse{Token: token_string}, nil
}

func (s *AccountsServiceServer) Create(ctx context.Context, request *accountspb.CreateRequest) (*accountspb.CreateResponse, error) {
	log := s.log.Named("CreateAccount")
	log.Debug("Create request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	ns, err := s.ns_ctrl.Get(ctx, request.Namespace)
	if err != nil {
		log.Debug("Error getting namespace", zap.Error(err), zap.String("namespace", request.Namespace))
		return nil, err
	}

	ok, access_lvl := graph.AccessLevel(ctx, s.db, requestor, ns.ID)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "No Access")
	} else if access_lvl < access.MGMT {
		return nil, status.Error(codes.PermissionDenied, "No Enough Rights")
	}

	acc, err := s.ctrl.Create(ctx, request.Title)
	if err != nil {
		log.Debug("Error creating account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating account")
	}
	res := &accountspb.CreateResponse{Uuid: acc.Key}

	if request.Access != nil && (*request.Access) < access_lvl {
		access_lvl = (*request.Access)
	}

	s.PostCreateActions(ctx, acc)

	col, _ := s.db.Collection(ctx, schema.NS2ACC)
	err = acc.JoinNamespace(ctx, col, ns, access_lvl, roles.OWNER)
	if err != nil {
		log.Debug("Error linking to namespace")
		return res, err
	}

	col, _ = s.db.Collection(ctx, schema.CREDENTIALS_EDGE_COL)
	cred, err := credentials.MakeCredentials(request.Auth, log)
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	err = s.ctrl.SetCredentials(ctx, acc, col, cred, roles.OWNER)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Update Account
// Supports updating Title and Data
//
//	Updating Data rules:
//		1. If Data is nil - it'll be skipped
//		2. If Data is not nil but has no keys - it'll be wiped
//		3. If value of one of the Data keys is nil - it'll be deleted from Data
func (s *AccountsServiceServer) Update(ctx context.Context, request *accountspb.Account) (*accountspb.UpdateResponse, error) {
	log := s.log.Named("UpdateAccount")
	log.Debug("Update request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	acc, err := graph.GetWithAccess[graph.Account](ctx, s.db, driver.NewDocumentID(schema.ACCOUNTS_COL, request.GetUuid()))
	if err != nil {
		log.Debug("Error getting account", zap.Any("error", err))
		return nil, status.Error(codes.NotFound, "Account not found")
	}
	if acc.Access == nil {
		log.Warn("Error Access is nil")
	}
	if acc.Access == nil || acc.Access.Level < accesspb.Level_ADMIN {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights to Account")
	}

	if acc.Access.Level < accesspb.Level_ROOT {
		request.SuspendConf = nil
		request.Suspended = nil
	}

	patch := make(map[string]interface{})

	if acc.Title != request.Title && request.Title != "" {
		log.Debug("Title patch detected")
		patch["title"] = request.Title
	}

	if request.Data == nil {
		log.Debug("Data patch is not present, skipping")
		goto patch
	}

	if len(request.Data.AsMap()) == 0 {
		log.Debug("Data patch is empty, wiping data")
		patch["data"] = nil
		goto patch
	}

	log.Debug("Merging data")
	patch["data"] = MergeMaps(acc.Data.AsMap(), request.Data.AsMap())

patch:
	if len(patch) == 0 {
		log.Debug("Resulting patch is empty")
		return nil, status.Error(codes.InvalidArgument, "Nothing changed")
	}

	log.Debug("Updating Account", zap.String("account", request.Uuid), zap.Any("patch", patch))
	err = s.ctrl.Update(ctx, acc, patch)
	if err != nil {
		log.Debug("Error updating account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating account")
	}

	return &accountspb.UpdateResponse{Result: true}, nil
}

func (s *AccountsServiceServer) EnsureRootExists(passwd string) error {
	return s.ctrl.EnsureRootExists(passwd)
}

func (s *AccountsServiceServer) SetCredentials(ctx context.Context, request *accountspb.SetCredentialsRequest) (*accountspb.SetCredentialsResponse, error) {
	log := s.log.Named("SetCredentials")
	log.Debug("Request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	acc, err := s.ctrl.Get(ctx, request.Account)
	if err != nil {
		log.Debug("Error getting account", zap.Any("error", err))
		return nil, status.Error(codes.NotFound, "Account not found")
	}

	if !graph.HasAccess(ctx, s.db, requestor, acc.ID, access.ADMIN) {
		return nil, status.Error(codes.PermissionDenied, "NoAccess")
	}

	auth := request.GetAuth()

	edge, _ := s.db.Collection(ctx, schema.ACC2CRED)
	old_cred_key, has_credentials := s.ctrl.GetCredentials(ctx, edge, acc, auth.Type)
	log.Debug("Checking if has credentials", zap.Bool("has_credentials", has_credentials), zap.Any("old_credentials", old_cred_key))

	cred, err := credentials.MakeCredentials(auth, log)
	if err != nil {
		log.Debug("Error creating new credentials", zap.String("type", auth.Type), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error creading new credentials")
	}

	if has_credentials {
		err = s.ctrl.UpdateCredentials(ctx, old_cred_key, cred)
	} else {
		err = s.ctrl.SetCredentials(ctx, acc, edge, cred, roles.OWNER)
	}

	if err != nil {
		log.Debug("Error updating/setting credentials", zap.String("type", auth.Type), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error updating/setting credentials")
	}

	return &accountspb.SetCredentialsResponse{Result: true}, nil
}

func (s *AccountsServiceServer) Delete(ctx context.Context, request *accountspb.DeleteRequest) (*accountspb.DeleteResponse, error) {
	log := s.log.Named("Delete")
	log.Debug("Request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	acc, err := s.ctrl.Get(ctx, request.Uuid)
	if err != nil {
		log.Debug("Error getting account", zap.Any("error", err))
		return nil, status.Error(codes.NotFound, "Account not found")
	}

	if !graph.HasAccess(ctx, s.db, requestor, acc.ID, access.ADMIN) {
		return nil, status.Error(codes.PermissionDenied, "NoAccess")
	}

	err = acc.Delete(ctx, s.db)
	if err != nil {
		log.Debug("Error deleting account and it's children", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error deleting account")
	}

	return &accountspb.DeleteResponse{Result: true}, nil
}
