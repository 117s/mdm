package initialize

import (
	"github.com/117s/x/config"
	"github.com/117s/x/logger"
)

type Config struct {
	Logger *logger.LoggerConfig `yaml:"logger"`
}

func LoadConfig(configPath *string) *Config {
	var configuration Config
	config.MustLoadConfig("MDM", configPath, &configuration)
	return &configuration
}
