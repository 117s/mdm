package service

import (
	"context"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/params"
	"github.com/117s/mdm/internal/schema"
)

func CreateDataModelDraft(ctx context.Context, p *params.CreateDataModelParams) (*schema.DataModel, error) {
	global.Log.Sugar().Debugf("create data model with prams: %+v", p)

	dm := &schema.DataModel{
		ID:         p.ID,
		Name:       p.Name,
		Identifier: p.Identifier,
		IsDraft:    true,
	}
	err := global.DB.WithContext(ctx).Model(schema.DataModel{}).Create(dm).Error
	if err != nil {
		global.Log.Sugar().Warnf("failed to create data model, err: %s", err)
		return nil, err
	}
	return dm, nil
}
