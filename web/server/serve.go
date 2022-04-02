package server

import (
	"context"
	"github.com/117s/mdm/web/initialize"
)

func Init(configfilepath *string) {
	initialize.LoadConfig(configfilepath)
}

func ServeAll(ctx context.Context, configfilepath *string) error {
	Init(configfilepath)
	return nil
}
