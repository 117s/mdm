package dto

type CreateDataModelDraftRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Name       string `json:"name" binding:"required"`
}

type CreateDataModelDraftResponse struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}
