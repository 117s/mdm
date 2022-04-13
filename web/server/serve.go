package server

import (
	"context"
	"fmt"
	"github.com/117s/mdm/internal/global"
)

func ServeAll(ctx context.Context, configfilepath *string) error {
	global.Init(configfilepath)

	return startHttpWriteServer()
}

func startHttpWriteServer() error {
	port := global.Config.System.Port
	addr := fmt.Sprintf(":%s", port)
	router := Routes()
	return router.Run(addr)
}
