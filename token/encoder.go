package leafToken

type (
	Encoder interface {
		EncodeToken(claims Claims) (string, error)
		CreateUserClaims(user UserLogin, exp int64) Claims
	}
)
