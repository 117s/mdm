package server

import (
	"github.com/117s/mdm/web/handler/data_model"
	"github.com/117s/mdm/web/handler/data_model_draft"
	"github.com/117s/mdm/web/handler/tenant"
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
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	v1 := router.Group("/api/v1")
	v1.Use(middleware.JWTAuthentication())

	v1.POST("/data-model-draft", data_model_draft.Create)
	v1.POST("/data-model-draft/:id/editing", data_model_draft.OnEditing)

	v1.GET("/data-model", data_model.Index)
	v1.GET("/data-model/:id", data_model.Show)
	v1.PUT("/data-model/:id", data_model.Publish)

	v1.POST("/tenants", tenant.Create)
	return router
}
