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
	"net"
	"strings"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/slntopp/nocloud/pkg/accounting/accountspb"
	"github.com/slntopp/nocloud/pkg/accounting/namespacespb"
	apipb "github.com/slntopp/nocloud/pkg/api/apipb"
	"github.com/slntopp/nocloud/pkg/health/healthpb"
	"github.com/slntopp/nocloud/pkg/nocloud"
	servicespb "github.com/slntopp/nocloud/pkg/services/proto"
	sppb "github.com/slntopp/nocloud/pkg/services_providers/proto"
)

var (
	log 			*zap.Logger

	healthHost 		string
	registryHost 	string
	servicesHost	string
	spRegistryHost  string

	SIGNING_KEY		[]byte
)

func resolveHost(addr string) string {
	host := strings.SplitN(addr, ":", 2)
	ips, err := net.LookupIP(host[0])
	if err != nil {
		log.Debug("Error resolving host", zap.String("host", host[0]), zap.Error(err))
		return addr
	}
	log.Debug("Resolved IPs", zap.Any("pool", ips))
	host[0] = ips[0].String()
	return strings.Join(host, ":")

}

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("CORS_ALLOWED", []string{"*"})

	viper.SetDefault("HEALTH_HOST", "health:8080")
	viper.SetDefault("REGISTRY_HOST", "accounts:8080")
	viper.SetDefault("SP_REGISTRY_HOST", "sp-registry:8080")
	viper.SetDefault("SERVICES_HOST", "services-registry:8080")
	
	viper.SetDefault("SIGNING_KEY", "seeeecreet")

	healthHost 		= resolveHost(viper.GetString("HEALTH_HOST"))
	registryHost 	= resolveHost(viper.GetString("REGISTRY_HOST"))
	servicesHost 	= resolveHost(viper.GetString("SERVICES_HOST"))
	spRegistryHost 	= resolveHost(viper.GetString("SP_REGISTRY_HOST"))

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

	log.Info("Connecting to ServicesProvidersService", zap.String("host", spRegistryHost))
	spRegistryConn, err := grpc.Dial(spRegistryHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	spRegistryClient := sppb.NewServicesProvidersServiceClient(spRegistryConn)

	log.Info("Connecting to ServicesService", zap.String("host", servicesHost))
	servicesConn, err := grpc.Dial(servicesHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	servicesClient := servicespb.NewServicesServiceClient(servicesConn)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen:", zap.Error(err))
	}

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(log),
			grpc.UnaryServerInterceptor(JWT_AUTH_INTERCEPTOR),
		)),
	)
	apipb.RegisterHealthServiceServer(s, &healthAPI{client: healthClient})

	apipb.RegisterAccountsServiceServer(s, &accountsAPI{client: accountsClient})
	apipb.RegisterNamespacesServiceServer(s, &namespacesAPI{client: namespacesClient})

	apipb.RegisterServicesProvidersServiceServer(s, &spRegistryAPI{client: spRegistryClient, log: log.Named("ServicesProvidersRegistry")})

	apipb.RegisterServicesServiceServer(s, &servicesAPI{client: servicesClient, log: log.Named("ServicesRegistry")})

	// Serve gRPC Server
	reflection.Register(s)
	log.Info("Serving gRPC on 0.0.0.0:8080", zap.Skip())
	log.Fatal("Error", zap.Error(s.Serve(lis)))

}
	