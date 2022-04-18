package dto

type CreateDataModelDraftRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Name       string `json:"name" binding:"required"`
	TenantID   string `json:"tenantId" binding:"required"`
}
