package server

import (
	"context"
	"fmt"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/initialize"
)

func ServeAll(ctx context.Context, configfilepath *string) error {
	initialize.Init(configfilepath)

	return startHttpWriteServer()
}

func startHttpWriteServer() error {
	port := global.Config.System.Port
	addr := fmt.Sprintf(":%s", port)
	router := Routes()
	global.Log.Debug("listen at http://127.0.0.1:" + port)
	return router.Run(addr)
}
