package leafJwt

import (
	"github.com/golang-jwt/jwt"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
)

type (
	encoderImpl struct {
		privateKey string
		alg        string
	}
)

func NewEncoder(privateKey string, alg string) leafToken.Encoder {
	return &encoderImpl{
		privateKey: privateKey,
		alg:        alg,
	}
}

func (i *encoderImpl) CreateUserClaims(user leafToken.UserLogin, exp int64) leafToken.Claims {
	return &claims{
		&jwt.StandardClaims{
			ExpiresAt: exp,
		},
		user,
	}
}

func (i *encoderImpl) EncodeToken(claims leafToken.Claims) (string, error) {
	privateKey := []byte(i.privateKey)

	newToken := jwt.New(jwt.GetSigningMethod(i.alg))
	newToken.Claims = claims

	signedToken, err := newToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
