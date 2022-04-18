package middleware

import (
	"github.com/117s/mdm/web/dto"
	"github.com/117s/mdm/web/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"net/http"
)

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := server.Server{}
		token, ok := s.BearerAuth(c.Request)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse("invalid token", "", ""))
			return
		}
		tid := utils.GetTenantID(token)
		c.Set("tid", tid)
	}
}
