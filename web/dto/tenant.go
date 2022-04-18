package dto

type CreateTenantRequest struct {
	Name string `json:"name" binding:"required"`
}
