package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token, err := global.OAuth2.ValidationBearerToken(c.Request)
		//if err != nil {
		//	global.Log.Sugar().Debugf("jwt verify failed, error: %s", err.Error())
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
		//		Message: "can not find jwt in Authorization header, or invalid jwt",
		//	})
		//	return
		//}
		//c.Set("token", token)
		//c.Set("client_id", token.GetClientID())
		//c.Set("user_id", token.GetUserID())
	}
}
