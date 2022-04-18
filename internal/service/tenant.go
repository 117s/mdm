package service

import (
	"context"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/schema"
)

func CreateTenant(ctx context.Context, tenant *schema.Tenant) (*schema.Tenant, error) {
	tx := global.DB.WithContext(ctx).Model(schema.Tenant{}).Create(tenant)
	return tenant, tx.Error
}
