package leafJwt

import leafToken "github.com/paulusrobin/leaf-utilities/token"

type (
	jwtImplementation struct {
		encoder leafToken.Encoder
		decoder leafToken.Decoder
	}
)

func (j *jwtImplementation) Encoder() leafToken.Encoder {
	return j.encoder
}

func (j *jwtImplementation) Decoder() leafToken.Decoder {
	return j.decoder
}

func (j *jwtImplementation) EncodeToken(claims leafToken.Claims) (string, error) {
	return j.encoder.EncodeToken(claims)
}

func (j *jwtImplementation) CreateUserClaims(user leafToken.UserLogin, exp int64) leafToken.Claims {
	return j.encoder.CreateUserClaims(user, exp)
}

func (j *jwtImplementation) DecodeToken(tokenString string) (leafToken.Claims, error) {
	return j.decoder.DecodeToken(tokenString)
}

func NewJWT(publicKey string, alg string) leafToken.Token {
	encoder := NewEncoder(publicKey, alg)
	decoder := NewDecoder(publicKey)
	return &jwtImplementation{encoder: encoder, decoder: decoder}
}
