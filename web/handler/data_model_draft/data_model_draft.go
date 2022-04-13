package data_model_draft

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDraft go docs
// @Summary Create a new `DataModel`
// @Description Create a new `DataModel`, start with a draft mode, need to be published before CRUD data
// @Tags DataModelDraft
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Failure      400   {object}  x.ErrorResponse
// @Failure      404   {object}  x.ErrorResponse
// @Failure      500   {object}  x.ErrorResponse
// @Router       /api/v1/data-model-draft [post]
func CreateDraft(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// OnEditing go docs
// @Summary marks a data model is being editing
// @Description marks this `DataModelDraft` is being edited by someone, thus can not be updated by anyone else.
// @Tags DataModelDraft
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Failure      400   {object}  x.ErrorResponse
// @Failure      404   {object}  x.ErrorResponse
// @Failure      500   {object}  x.ErrorResponse
// @Router       /api/v1/data-model-draft/{id}/editing [post]
func OnEditing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
