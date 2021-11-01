/*
Copyright © 2021 Nikita Ivanovski info@slnt-opp.xyz

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
package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"

	"github.com/slntopp/nocloud/pkg/accounting/accountspb"
	"github.com/slntopp/nocloud/pkg/accounting/namespacespb"
	apipb "github.com/slntopp/nocloud/pkg/api/apipb"
	"github.com/slntopp/nocloud/pkg/health/healthpb"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	log 			*zap.Logger

	healthHost 		string
	registryHost 	string

	SIGNING_KEY		[]byte
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("HEALTH_HOST", "health:8080")
	viper.SetDefault("REGISTRY_HOST", "accounts:8080")
	
	viper.SetDefault("SIGNING_KEY", "seeeecreet")

	healthHost 		= viper.GetString("HEALTH_HOST")
	registryHost 	= viper.GetString("REGISTRY_HOST")

	SIGNING_KEY 	= []byte(viper.GetString("SIGNING_KEY"))
}

func main() {
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Connecting to HealthService", zap.String("host", healthHost))
	healthConn, err := grpc.Dial(healthHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	healthClient := healthpb.NewHealthServiceClient(healthConn)

	log.Info("Connecting to AccountsService", zap.String("host", registryHost))
	registryConn, err := grpc.Dial(registryHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	accountsClient := accountspb.NewAccountsServiceClient(registryConn)
	namespacesClient := namespacespb.NewNamespacesServiceClient(registryConn)
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen:", zap.Error(err))
	}

	// Create a gRPC server object
	s := grpc.NewServer(grpc.UnaryInterceptor(JWT_AUTH_INTERCEPTOR),)
	// Attach the Greeter service to the server
	apipb.RegisterHealthServiceServer(s, &healthAPI{client: healthClient})
	apipb.RegisterAccountsServiceServer(s, &accountsAPI{client: accountsClient})
	apipb.RegisterNamespacesServiceServer(s, &namespacesAPI{client: namespacesClient})
	// Serve gRPC Server
	log.Info("Serving gRPC on 0.0.0.0:8080", zap.Skip())
	go func() {
		log.Fatal("Error", zap.Error(s.Serve(lis)))
	}()

	// Set up REST API server
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("Failed to dial server:", zap.Error(err))
	}

	gwmux := runtime.NewServeMux()
	err = apipb.RegisterHealthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register HealthService gateway", zap.Error(err))
	}
	err = apipb.RegisterAccountsServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register AccountsService gateway", zap.Error(err))
	}
	err = apipb.RegisterNamespacesServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register NamespacesService gateway", zap.Error(err))
	}
	gwServer := &http.Server{
		Addr:    ":8000",
		Handler: gwmux,
	}

	log.Info("Serving gRPC-Gateway on http://0.0.0.0:8000")
	log.Fatal("Failed to Listen and Serve Gateway-Server", zap.Error(gwServer.ListenAndServe()))
}