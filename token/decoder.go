package leafToken

type (
	Decoder interface {
		DecodeToken(tokenString string) (Claims, error)
	}
)
