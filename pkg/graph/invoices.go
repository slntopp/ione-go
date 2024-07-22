package graph

import (
	"context"
	"errors"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/mitchellh/mapstructure"
	pb "github.com/slntopp/nocloud-proto/billing"
	"github.com/slntopp/nocloud/pkg/graph/migrations"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"
	"strings"
	"time"
)

type Invoice struct {
	*pb.Invoice    `json:"invoice"`
	NumericNumber  int    `json:"numeric_number"`
	NumberTemplate string `json:"number_template"`
	driver.DocumentMeta
}

type InvoicesController struct {
	col driver.Collection // Billing Plans collection

	log *zap.Logger
}

func NewInvoicesController(logger *zap.Logger, db driver.Database) InvoicesController {
	ctx := context.TODO()
	log := logger.Named("InvoicesController")
	col := GetEnsureCollection(log, ctx, db, schema.INVOICES_COL)

	migrations.UpdateNumericCurrencyToDynamic(log, col)

	return InvoicesController{
		log: log, col: col,
	}
}

func (ctrl *InvoicesController) ParseNumberIntoTemplate(template string, number int, date time.Time) string {
	year := date.Year()
	month := int(date.Month())
	day := date.Day()
	template = strings.Replace(template, "{YEAR}", fmt.Sprintf("%d", year), -1)
	template = strings.Replace(template, "{MONTH}", fmt.Sprintf("%d", month), -1)
	template = strings.Replace(template, "{DAY}", fmt.Sprintf("%d", day), -1)
	template = strings.Replace(template, "{NUMBER}", fmt.Sprintf("%d", number), -1)
	return template
}

func (ctrl *InvoicesController) Create(ctx context.Context, tx *Invoice) (*Invoice, error) {
	if tx.GetAccount() == "" {
		return nil, errors.New("account is required")
	}
	if tx.Total == 0 {
		return nil, errors.New("total is required")
	}
	meta, err := ctrl.col.CreateDocument(ctx, tx)
	if err != nil {
		ctrl.log.Error("Failed to create transaction", zap.Error(err))
		return nil, err
	}
	tx.Uuid = meta.Key
	return tx, nil
}

func (ctrl *InvoicesController) Get(ctx context.Context, uuid string) (*Invoice, error) {
	var tx = &Invoice{}
	tx.Invoice = &pb.Invoice{}
	result := map[string]interface{}{}
	meta, err := ctrl.col.ReadDocument(ctx, uuid, &result)
	if err != nil {
		ctrl.log.Error("Failed to read invoice", zap.Error(err))
		return nil, err
	}
	err = mapstructure.Decode(result["invoice"], tx.Invoice)
	tx.Uuid = meta.Key
	return tx, err
}

func (ctrl *InvoicesController) Update(ctx context.Context, tx *Invoice) (*Invoice, error) {
	_, err := ctrl.col.UpdateDocument(ctx, tx.GetUuid(), tx)
	if err != nil {
		ctrl.log.Error("Failed to update invoice", zap.Error(err))
		return nil, err
	}
	return tx, nil
}

func (ctrl *InvoicesController) Patch(ctx context.Context, id string, patch map[string]interface{}) error {
	_, err := ctrl.col.UpdateDocument(ctx, id, patch)
	return err
}
