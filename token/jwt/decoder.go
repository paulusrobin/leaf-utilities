package leafJwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
)

type (
	implementation struct {
		publicKey string
	}
)

func NewDecoder(publicKey string) leafToken.Decoder {
	return &implementation{
		publicKey: publicKey,
	}
}

func (i *implementation) DecodeToken(tokenString string) (leafToken.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(i.publicKey), nil
	})
	if err != nil || tokenClaims == nil || tokenClaims.Claims == nil {
		return nil, fmt.Errorf("unclaimed token")
	}

	claims := tokenClaims.Claims.(*claims)
	return claims, nil
}
