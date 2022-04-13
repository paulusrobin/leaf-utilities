package leafJwt

import (
	"github.com/golang-jwt/jwt"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
	"time"
)

type (
	claims struct {
		*jwt.StandardClaims
		UserLogin leafToken.UserLogin `json:"user"`
	}
)

func (c *claims) User() leafToken.UserLogin {
	return c.UserLogin
}

func (c *claims) Valid() error {
	jwt.TimeFunc = time.Now

	return c.StandardClaims.Valid()
}
