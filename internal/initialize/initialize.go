package initialize

import (
	"github.com/117s/mdm/internal/config"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/x/logger"
	"github.com/117s/x/viper"
	"github.com/go-playground/validator/v10"
)

func Init(configFilepath *string) {
	var configuration config.Config
	viper.MustLoadConfig("MDM", configFilepath, &configuration)
	global.Config = &configuration

	l := logger.MustGetLogger(configuration.Logger)
	global.Log = l

	global.Validator = validator.New()

	// connect to metadata server, panic if any err occurs
	db := Gorm()
	if db == nil {
		panic("failed to connect to metadata database")
	}
	global.DB = db

	RegisterTables(db)
}
