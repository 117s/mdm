package dto

type ErrorResponse struct {
	Message string
	Code    string
	Details interface{}
}

type PaginationResult struct {
	Results  interface{}        `json:"results"`
	Metadata PaginationMetadata `json:"metadata"`
}

type PaginationMetadata struct {
	Total   int  `json:"total"`
	Skip    int  `json:"skip"`
	Limit   int  `json:"limit"`
	HasNext bool `json:"hasNext"`
}

type PageQueryParams struct {
	Skip  int `form:"skip"`
	Limit int `form:"limit"`
}

type CommonQuery struct {
	PageQueryParams
	SortBy *string `form:"sortby"`
	Filter *string `form:"filter"`
}

func NewErrorResponse(message, code string, details interface{}) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Code:    code,
		Details: details,
	}
}

func ErrInvalidBody(details interface{}) ErrorResponse {
	return ErrorResponse{
		Message: "invalid request body",
		Code:    "common-1",
		Details: details,
	}
}

func ErrInvalidQuery(details interface{}) ErrorResponse {
	return ErrorResponse{
		Message: "invalid request query params",
		Code:    "common-2",
		Details: details,
	}
}
func ErrUnauthorized(details interface{}) ErrorResponse {
	return ErrorResponse{
		Message: "unauthorized",
		Code:    "common-2",
		Details: details,
	}
}
