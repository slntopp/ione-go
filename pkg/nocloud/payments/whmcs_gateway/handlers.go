package whmcs_gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type ContextKey string

const GatewayCallback = ContextKey("payment-gateway-callback")

func (g *WhmcsGateway) invoiceCreatedHandler(ctx context.Context, log *zap.Logger, data InvoiceCreated) error {
	whmcsInv, err := g.GetInvoice(context.Background(), data.InvoiceId)
	if err != nil {
		return err
	}

	_, err = g.GetAccountByWhmcsId(whmcsInv.UserId)
	if err != nil {
		return err
	}

	// TODO

	return nil
}

func (g *WhmcsGateway) invoiceDeletedHandler(ctx context.Context, log *zap.Logger, data InvoiceDeleted) error {
	return fmt.Errorf("not implemented invoicedeleted logic")
}

func unmarshal[T any](b []byte) (T, error) {
	var res T
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (g *WhmcsGateway) handleWhmcsEvent(log *zap.Logger, body []byte) {
	log = log.Named("WhmcsHandler")

	resp := struct {
		Event string `json:"event"`
	}{}
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Error("Error decoding request", zap.Error(err), zap.String("body", string(body)))
		return
	}

	log.Info("Event received", zap.String("event", resp.Event))
	log = log.With(zap.String("event", resp.Event))

	// Now process only PAID and REFUNDED events. Only these events are valuable due to after-paid actions
	ctx := context.WithValue(context.Background(), GatewayCallback, true)
	var innerErr error
	switch resp.Event {
	case "InvoicePaid":
		data, err := unmarshal[InvoicePaid](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		innerErr = g.syncWhmcsInvoice(ctx, data.InvoiceId)
	case "InvoiceModified":
		_, err := unmarshal[InvoiceModified](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		//innerErr = g.syncWhmcsInvoice(ctx, data.InvoiceId)
	case "InvoiceCancelled":
		_, err := unmarshal[InvoiceCancelled](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		//innerErr = g.syncWhmcsInvoice(ctx, data.InvoiceId)
	case "InvoiceRefunded":
		data, err := unmarshal[InvoiceRefunded](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		innerErr = g.syncWhmcsInvoice(ctx, data.InvoiceId)
	case "InvoiceUnpaid":
		_, err := unmarshal[InvoiceUnpaid](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		//innerErr = g.syncWhmcsInvoice(ctx, data.InvoiceId)
	case "UpdateInvoiceTotal":
		_, err := unmarshal[UpdateInvoiceTotal](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		//innerErr = g.syncWhmcsInvoice(ctx, data.InvoiceId)
	case "InvoiceCreated":
		_, err := unmarshal[InvoiceCreated](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		//innerErr = g.invoiceCreatedHandler(ctx, log, data)
	case "InvoiceDeleted":
		_, err := unmarshal[InvoiceDeleted](body)
		if err != nil {
			log.Error("Error decoding request", zap.Error(err))
			return
		}
		//innerErr = g.invoiceDeletedHandler(ctx, log, data)
	default:
		log.Error("Unknown event", zap.String("event", resp.Event))
		return
	}
	if innerErr != nil {
		log.Error("Error handling event", zap.Error(innerErr))
	}
}
