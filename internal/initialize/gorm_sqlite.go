package initialize

import (
	"github.com/117s/mdm/internal/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormSqlite 初始化 Sqlite 数据库
func GormSqlite() *gorm.DB {
	p := global.Config.Sqlite
	db, err := gorm.Open(sqlite.Open(p.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil
	}
	return db
}
