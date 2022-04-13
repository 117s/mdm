package config

import "github.com/117s/x/logger"

type Config struct {
	System *System              `mapstructure:"system" json:"system" yaml:"system"`
	Logger *logger.LoggerConfig `mapstructure:"logger" json:"logger" yaml:"logger"`
}
