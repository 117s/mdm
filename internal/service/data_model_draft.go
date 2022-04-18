package service

import (
	"context"
	"errors"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/params"
	"github.com/117s/mdm/internal/schema"
	"gorm.io/gorm"
)

func CreateDataModelDraft(ctx context.Context, p *params.CreateDataModelParams) (*schema.DataModel, error) {
	global.Log.Sugar().Debugf("create data model with prams: %+v", p)

	// check existence of this data model and tenant
	_, exists, err := GetDataModelByIdentifier(ctx, p.TenantID, p.Identifier)
	if exists {
		return nil, errors.New("data model with identifier " + p.Identifier + " already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	dm := &schema.DataModel{
		ID:         p.ID,
		Name:       p.Name,
		Identifier: p.Identifier,
		IsDraft:    true,
		TenantID:   p.TenantID,
	}
	err = global.DB.WithContext(ctx).Model(schema.DataModel{}).Create(dm).Error
	if err != nil {
		global.Log.Sugar().Warnf("failed to create data model, err: %s", err)
		return nil, err
	}
	return dm, nil
}
