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
package graph

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
	pb "github.com/slntopp/nocloud-proto/billing"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"
)

type Transaction struct {
	*pb.Transaction
	driver.DocumentMeta
}

type TransactionsController struct {
	col driver.Collection // Billing Plans collection

	log *zap.Logger
}

func NewTransactionsController(logger *zap.Logger, db driver.Database) TransactionsController {
	ctx := context.TODO()
	log := logger.Named("TransactionsController")
	col := GetEnsureCollection(log, ctx, db, schema.TRANSACTIONS_COL)
	return TransactionsController{
		log: log, col: col,
	}
}

func (ctrl *TransactionsController) Create(ctx context.Context, tx *pb.Transaction) (*Transaction, error) {
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
	return &Transaction{tx, meta}, nil
}

func (ctrl *TransactionsController) Get(ctx context.Context, uuid string) (*pb.Transaction, error) {
	var tx pb.Transaction
	_, err := ctrl.col.ReadDocument(ctx, uuid, &tx)
	if err != nil {
		ctrl.log.Error("Failed to read transaction", zap.Error(err))
		return nil, err
	}
	return &tx, nil
}
