package initialize

import (
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/schema"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm initialize the metadata database and store gorm DB with a global ref
func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormPgSql()
	}
}

// RegisterTables register metadata tables
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		schema.DataModel{},
		schema.Property{},
	)
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		panic("failed to register tables")
	}
	global.Log.Info("successfully registered tables")
}
