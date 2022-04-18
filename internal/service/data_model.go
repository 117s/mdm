package service

import (
	"context"
	"errors"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/schema"
	"github.com/117s/mdm/web/dto"
	"gorm.io/gorm"
)

// GetDataModels find data models from db with given query params
func GetDataModels(ctx context.Context, tenantId string, query dto.CommonQuery) ([]schema.DataModel, int64, error) {
	global.Log.Sugar().Debugf("get data model list with query: %+v", query)

	// TODO: support order and filter
	q := global.DB.WithContext(ctx).Model(schema.DataModel{}).Where(schema.DataModel{TenantID: tenantId})
	var cnt int64
	tx := q.Count(&cnt)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	var models []schema.DataModel
	tx = q.Offset(query.Skip).Limit(query.Limit).Find(&models)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	return models, cnt, nil
}

func GetDataModelByID(ctx context.Context, id string) (*schema.DataModel, error) {
	global.Log.Sugar().Debugf("get data model %s", id)

	var dm schema.DataModel
	tx := global.DB.WithContext(ctx).Model(schema.DataModel{}).Preload("Properties").Take(&dm, schema.DataModel{ID: id})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dm, nil
}

func GetDataModelByIdentifier(ctx context.Context, tenantID, identifier string) (*schema.DataModel, bool, error) {
	global.Log.Sugar().Debugf("get data model by identifier:%s, tenantId: %s", identifier, tenantID)

	var dm schema.DataModel
	tx := global.DB.WithContext(ctx).Model(schema.DataModel{}).
		Where(schema.DataModel{Identifier: identifier, TenantID: tenantID}).
		Take(&dm)
	exists := !errors.Is(tx.Error, gorm.ErrRecordNotFound)

	return &dm, exists, tx.Error
}
