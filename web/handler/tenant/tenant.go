package tenant

import (
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/schema"
	"github.com/117s/mdm/internal/service"
	"github.com/117s/mdm/internal/utils"
	"github.com/117s/mdm/web/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create go docs
// @Summary      Create a new Tenant
// @Description  Create a new `Tenant`
// @Tags         Tenant
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        data  body      dto.CreateTenantRequest  true  "create new tenant request"
// @Success      200   {object}  schema.Tenant
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      404   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /api/v1/tenant [post]
func Create(c *gin.Context) {
	var req dto.CreateTenantRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewRequestBodyErr(err.Error()))
		return
	}
	global.Log.Sugar().Debugf("create tenant with %+v", req)

	tenant := &schema.Tenant{
		ID:   utils.NewID(),
		Name: req.Name,
	}
	res, err := service.CreateTenant(c, tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "", ""))
		return
	}

	c.JSON(http.StatusOK, res)
}
