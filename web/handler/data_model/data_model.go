package data_model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish go docs
// @Summary Publish a `DataModel` draft, after publish, this DataModel will really take effects on database.
// @Tags DataModel
func Publish(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
