package global

import (
	"github.com/117s/mdm/internal/config"
	"github.com/117s/x/logger"
	"github.com/117s/x/viper"
	v10 "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	Log       *zap.Logger
	Config    *config.Config
	Validator *v10.Validate
)

func Init(configFilepath *string) {
	var configuration config.Config
	viper.MustLoadConfig("MDM", configFilepath, &configuration)
	Config = &configuration

	l := logger.MustGetLogger(configuration.Logger)
	Log = l

	Validator = v10.New()
}
