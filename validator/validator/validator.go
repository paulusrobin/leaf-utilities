package leafValidator

type (
	Validator interface {
		Validate(interface{}) error
	}
)
