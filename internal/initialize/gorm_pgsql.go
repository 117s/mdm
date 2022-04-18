package initialize

import (
	"github.com/117s/mdm/internal/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql() *gorm.DB {
	p := global.Config.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
