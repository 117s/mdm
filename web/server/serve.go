package server

import (
	"context"
	"github.com/117s/mdm/internal/global"
)

func ServeAll(ctx context.Context, configfilepath *string) error {
	global.Init(configfilepath)
	return nil
}
