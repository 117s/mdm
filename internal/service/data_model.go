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
func GetDataModels(ctx context.Context, query dto.CommonQuery) ([]schema.DataModel, error) {
	global.Log.Sugar().Debugf("get data model list with query: %+v", query)

	var models []schema.DataModel
	// TODO: support order and filter
	// TODO: filter by tenant id
	tx := global.DB.WithContext(ctx).Model(schema.DataModel{}).
		Offset(query.Skip).
		Limit(query.Limit).
		Find(&models)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return models, nil
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
