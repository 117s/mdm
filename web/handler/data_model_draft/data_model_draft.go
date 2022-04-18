package data_model_draft

import (
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/params"
	"github.com/117s/mdm/internal/service"
	"github.com/117s/mdm/internal/utils"
	"github.com/117s/mdm/web/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDraft go docs
// @Summary      Create a new `DataModel`
// @Description  Create a new `DataModel`, start with a draft mode, need to be published before CRUD data
// @Tags         DataModelDraft
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        data  body      dto.CreateDataModelDraftRequest  true  "create data model request"
// @Success      200   {object}  dto.CreateDataModelDraftResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/v1/data-model-draft [post]
func CreateDraft(c *gin.Context) {
	var req dto.CreateDataModelDraftRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewRequestBodyErr(err.Error()))
		return
	}
	global.Log.Sugar().Debugf("create data model draft with %+v", req)

	dm := &params.CreateDataModelParams{
		ID:         utils.NewID(),
		Identifier: req.Identifier,
		Name:       req.Name,
	}
	res, err := service.CreateDataModelDraft(c, dm)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "", ""))
		return
	}

	c.JSON(http.StatusOK, res)
}

// OnEditing go docs
// @Summary      marks a data model is being editing
// @Description  marks this `DataModelDraft` is being edited by someone, thus can not be updated by anyone else.
// @Tags         DataModelDraft
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      404   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /api/v1/data-model-draft/{id}/editing [post]
func OnEditing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
