package graph

import (
	"context"
	"testing"

	pb "github.com/slntopp/nocloud/pkg/billing/proto"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/auth"
	"github.com/slntopp/nocloud/pkg/nocloud/connectdb"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func TestConvert(t *testing.T) {
	viper.SetDefault("DB_CRED", "root:openSesame")
	viper.SetDefault("DB_HOST", "localhost:8529")
	viper.SetDefault("SIGNING_KEY", "seeeecreet")

	arangodbHost = viper.GetString("DB_HOST")
	arangodbCred = viper.GetString("DB_CRED")

	log := nocloud.NewLogger()
	db := connectdb.MakeDBConnection(log, arangodbHost, arangodbCred)

	SIGNING_KEY := []byte(viper.GetString("SIGNING_KEY"))

	auth.SetContext(log, SIGNING_KEY)

	token, err := auth.MakeToken(schema.ROOT_ACCOUNT_KEY)
	if err != nil {
		log.Fatal("Can't generate token", zap.Error(err))
	}

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "bearer "+token)

	c := NewCurrencyController(log, db)
	list, err := c.Get(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(list) != len(currencies) {
		t.Error("Default currencies haven't been created")
	}

	testRate := 2.0
	c.CreateExchangeRate(ctx, pb.Currency_USD, pb.Currency_BYN, testRate)

	rate, err := c.GetExchangeRate(ctx, pb.Currency_USD, pb.Currency_BYN)
	if err != nil {
		t.Error(err)
	}
	if rate != testRate {
		t.Error("Wrong exchange rate")
	}

	amount, err := c.Convert(ctx, pb.Currency_USD, pb.Currency_BYN, 1.0)
	if err != nil {
		t.Error(err)
	}
	wanted := 2.0
	if amount != wanted {
		t.Errorf("Wrong conversion. Wanted %f. Got = %f", wanted, amount)
	}

	err = c.DeleteExchangeRate(ctx, pb.Currency_USD, pb.Currency_BYN)
	if err != nil {
		t.Error(err)
	}
}
