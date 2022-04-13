package server

import (
	"github.com/117s/mdm/web/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/117s/mdm/docs"
)

func Routes() *gin.Engine {
	router := gin.Default()

	// middleware
	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
