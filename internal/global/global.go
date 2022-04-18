package global

import (
	"github.com/117s/mdm/internal/config"
	v10 "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log       *zap.Logger
	DB        *gorm.DB
	Config    *config.Config
	Validator *v10.Validate
)
