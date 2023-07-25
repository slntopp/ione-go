package graph

import (
	"context"
	"github.com/arangodb/go-driver"
	sppb "github.com/slntopp/nocloud-proto/services_providers"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"
)

type ShowcasesController struct {
	log *zap.Logger
	col driver.Collection

	db driver.Database
}

func NewShowcasesController(logger *zap.Logger, db driver.Database) *ShowcasesController {
	ctx := context.Background()
	log := logger.Named("ShowcasesController")
	log.Debug("New Showcases Controller Creating")

	col := GetEnsureCollection(log, ctx, db, schema.SHOWCASES_COL)

	return &ShowcasesController{
		log: log,
		col: col,
		db:  db,
	}
}

func (ctrl *ShowcasesController) Create(ctx context.Context, showcase *sppb.Showcase) (*sppb.Showcase, error) {
	ctrl.log.Debug("Creating Document for Showcase", zap.Any("config", showcase))

	meta, err := ctrl.col.CreateDocument(ctx, showcase)
	if err != nil {
		return nil, err
	}

	showcase.Uuid = meta.Key

	_, err = ctrl.col.UpdateDocument(ctx, showcase.GetUuid(), showcase)
	if err != nil {
		return nil, err
	}

	return showcase, err
}

func (ctrl *ShowcasesController) Update(ctx context.Context, showcase *sppb.Showcase) (*sppb.Showcase, error) {
	ctrl.log.Debug("Updating ServicesProvider", zap.Any("sp", showcase))

	meta, err := ctrl.col.ReplaceDocument(ctx, showcase.GetUuid(), showcase)
	ctrl.log.Debug("ReplaceDocument.Result", zap.Any("meta", meta), zap.Error(err))
	return showcase, err
}

const listQuery = `
FOR s IN @@showcases
	RETURN s
`

func (ctrl *ShowcasesController) List(ctx context.Context) ([]*sppb.Showcase, error) {
	ctrl.log.Debug("Getting Showcases")

	c, err := ctrl.col.Database().Query(ctx, listQuery, map[string]interface{}{
		"@showcases": schema.SHOWCASES_COL,
	})
	if err != nil {
		return nil, err
	}

	defer c.Close()
	var r []*sppb.Showcase
	for c.HasMore() {
		var s sppb.Showcase

		_, err := c.ReadDocument(ctx, &s)

		if err != nil {
			return nil, err
		}

		ctrl.log.Debug("Got document", zap.Any("service_provider", &s))
		r = append(r, &s)
	}

	return r, nil
}

func (ctrl *ShowcasesController) Get(ctx context.Context, uuid string) (*sppb.Showcase, error) {
	ctrl.log.Debug("Getting Showcase", zap.Any("uuid", uuid))

	var showcase sppb.Showcase
	meta, err := ctrl.col.ReadDocument(ctx, uuid, &showcase)
	ctrl.log.Debug("ReplaceDocument.Result", zap.Any("meta", meta), zap.Error(err))
	return &showcase, err
}

func (ctrl *ShowcasesController) Delete(ctx context.Context, uuid string) error {
	ctrl.log.Debug("Deleting Showcase", zap.Any("uuid", uuid))

	meta, err := ctrl.col.RemoveDocument(ctx, uuid)
	ctrl.log.Debug("ReplaceDocument.Result", zap.Any("meta", meta), zap.Error(err))
	return err
}
