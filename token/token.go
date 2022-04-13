package leafToken

type (
	Token interface {
		Encoder
		Decoder
		Encoder() Encoder
		Decoder() Decoder
	}
)
