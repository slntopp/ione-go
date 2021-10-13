package graph

import (
	"context"

	"github.com/arangodb/go-driver"
)

var (
	ACCOUNTS_COL = "Accounts"
)

type Account struct {
	Title string `json:"title"`
	driver.DocumentMeta
}

type AccountsController struct {
	col driver.Collection
}

// func Create(title string) (string, error) {}

func (ctrl *AccountsController) Get(ctx context.Context, id string) (Account, error) {
	var r Account
	_, err := ctrl.col.ReadDocument(nil, id, &r)
	return r, err
}

func (ctrl *AccountsController) Exists(ctx context.Context, id string) (bool, error) {
	return ctrl.col.DocumentExists(nil, id)
}

func (ctrl *AccountsController) Create(ctx context.Context, title string) (Account, error) {
	acc := Account{
		Title: title,
	}
	meta, err := ctrl.col.CreateDocument(ctx, acc)
	acc.DocumentMeta = meta
	return acc, err
}

// Grant account access to namespace
func (acc *Account) LinkNamespace(ctx context.Context, col driver.Collection, ns Namespace, level int8) (error) {
	_, err := col.CreateDocument(ctx, Access{
		From: acc.ID,
		To: ns.ID,
		Level: level,
		DocumentMeta: driver.DocumentMeta {
			Key: acc.Key + "-" + ns.Key,
		},
	})
	return err
}