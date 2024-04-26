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
package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"

	cc "github.com/slntopp/nocloud-proto/billing/billingconnect"
	billing "github.com/slntopp/nocloud/pkg/billing"
	"github.com/slntopp/nocloud/pkg/nocloud"
	auth "github.com/slntopp/nocloud/pkg/nocloud/connect_auth"
	"github.com/slntopp/nocloud/pkg/nocloud/connectdb"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"

	"connectrpc.com/grpchealth"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

var (
	port string
	log  *zap.Logger

	RabbitMQConn string
	redisHost    string
	arangodbHost string
	arangodbCred string
	SIGNING_KEY  []byte
)

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("PORT", "8000")

	viper.SetDefault("REDIS_HOST", "redis:6379")
	viper.SetDefault("DB_HOST", "db:8529")
	viper.SetDefault("DB_CRED", "root:openSesame")
	viper.SetDefault("DRIVERS", "")
	viper.SetDefault("EXTENTION_SERVERS", "")
	viper.SetDefault("SIGNING_KEY", "seeeecreet")

	port = viper.GetString("PORT")

	arangodbHost = viper.GetString("DB_HOST")
	arangodbCred = viper.GetString("DB_CRED")
	redisHost = viper.GetString("REDIS_HOST")
	SIGNING_KEY = []byte(viper.GetString("SIGNING_KEY"))

	viper.SetDefault("RABBITMQ_CONN", "amqp://nocloud:secret@rabbitmq:5672/")
	RabbitMQConn = viper.GetString("RABBITMQ_CONN")
}

func main() {
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Setting up DB Connection")
	db := connectdb.MakeDBConnection(log, arangodbHost, arangodbCred)
	log.Info("DB connection established")

	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost,
		DB:   0,
	})

	authInterceptor := auth.NewInterceptor(log, rdb, SIGNING_KEY)
	interceptors := connect.WithInterceptors(authInterceptor)

	router := mux.NewRouter()
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("Request", zap.String("method", r.Method), zap.String("path", r.URL.Path))
			h.ServeHTTP(w, r)
		})
	})

	conn, err := amqp.Dial(RabbitMQConn)
	if err != nil {
		log.Fatal("failed to connect to RabbitMQ", zap.Error(err))
	}
	defer conn.Close()

	server := billing.NewBillingServiceServer(log, db, conn)
	currencies := billing.NewCurrencyServiceServer(log, db)
	log.Info("Starting Currencies Service")

	token, err := authInterceptor.MakeToken(schema.ROOT_ACCOUNT_KEY)
	if err != nil {
		log.Fatal("Can't generate token", zap.Error(err))
	}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "bearer "+token)

	log.Info("Starting Transaction Generator-Processor")
	go server.GenTransactionsRoutine(ctx)

	log.Info("Starting Account Suspension Routine")
	go server.SuspendAccountsRoutine(ctx)

	log.Info("Registering BillingService Server")
	path, handler := cc.NewBillingServiceHandler(server, interceptors)
	router.PathPrefix(path).Handler(handler)

	records := billing.NewRecordsServiceServer(log, conn, db)
	log.Info("Starting Records Consumer")
	go records.Consume(ctx)
	go server.Consume(ctx) // Expiring records consumer

	log.Info("Registering CurrencyService Server")
	path, handler = cc.NewCurrencyServiceHandler(currencies, interceptors)
	router.PathPrefix(path).Handler(handler)

	addons := billing.NewAddonsServer(log, db)
	log.Info("Registering AddonsService Server")
	path, handler = cc.NewAddonsServiceHandler(addons, interceptors)
	router.PathPrefix(path).Handler(handler)

	descriptions := billing.NewDescriptionsServer(log, db)
	log.Info("Registering descriptionsService Server")
	path, handler = cc.NewDescriptionsServiceHandler(descriptions, interceptors)
	router.PathPrefix(path).Handler(handler)

	checker := grpchealth.NewStaticChecker()
	path, handler = grpchealth.NewHandler(checker)
	router.PathPrefix(path).Handler(handler)

	host := fmt.Sprintf("0.0.0.0:%s", port)

	handler = cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowedMethods:      []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:      []string{"*", "Connect-Protocol-Version"},
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
	}).Handler(h2c.NewHandler(router, &http2.Server{}))

	log.Info("Serving", zap.String("host", host))
	err = http.ListenAndServe(host, handler)
	if err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}

}
