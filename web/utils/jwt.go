package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	jwt.StandardClaims
	TenantID string
}

var mySigningKey = []byte("AllYourBase")

func GetTenantID(ss string) string {
	token, err := jwt.ParseWithClaims(ss, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if cliaims, ok := token.Claims.(*TokenClaims); ok {
		return cliaims.TenantID
	} else {
		fmt.Println(err)
	}
	return ""
}

// GenerateToken is only used for testing purpose!!
func GenerateToken(tenantId string) string {
	claims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
		TenantID: tenantId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Println(ss)
	return ss
}
