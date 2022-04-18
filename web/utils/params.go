package utils

import (
	"github.com/117s/mdm/web/dto"
	"github.com/gin-gonic/gin"
)

func ParsePaginationQuery(c *gin.Context) (dto.CommonQuery, error) {
	skip := c.Query("skip")
	limit := c.Query("limit")
	// TODO: make default skip and limit configurable
	defaultLimit := 20
	defaultSkip := 0

	var params dto.CommonQuery
	err := c.BindQuery(&params)
	if err != nil {
		return params, err
	}

	if skip == "" || params.Skip < 0 {
		params.Skip = defaultSkip
	}
	if limit == "" || params.Limit <= 0 {
		params.Limit = defaultLimit
	}

	return params, nil
}

func ToPaginationResult(results interface{}, resultLength, cnt, skip, limit int) dto.PaginationResult {
	return dto.PaginationResult{
		Results: results,
		Metadata: dto.PaginationMetadata{
			Total:   cnt,
			Skip:    skip,
			Limit:   limit,
			HasNext: (resultLength + skip) < cnt,
		},
	}
}
