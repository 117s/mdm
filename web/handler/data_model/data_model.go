package data_model

import (
	"github.com/117s/mdm/internal/service"
	"github.com/117s/mdm/web/dto"
	"github.com/117s/mdm/web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish go docs
// @Summary  Publish a `DataModel` draft, after publish, this DataModel will really take effects on database.
// @Tags     DataModel
func Publish(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Index go docs
// @Summary      List all DataModel
// @Description  List all `DataModel`s
// @Tags         DataModel
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  dto.PaginationResult
// @Success      200  {object}  dto.PaginationResult{results=[]schema.DataModel}
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/v1/data-model [get]
func Index(c *gin.Context) {
	query, err := utils.ParsePaginationQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewRequestQueryErr(err.Error()))
		return
	}

	res, err := service.GetDataModels(c, query)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "", ""))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Show go docs
// @Summary      Get a DataModel details
// @Description  Get a DataModel details
// @Tags         DataModel
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  schema.DataModel
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/v1/data-model/{id} [get]
func Show(c *gin.Context) {
	modelId := c.Param("id")
	model, err := service.GetDataModelByID(c, modelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "", ""))
		return
	}
	c.JSON(http.StatusOK, model)
}
