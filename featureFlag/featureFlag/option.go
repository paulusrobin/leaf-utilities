package featureFlag

import "time"

type (
	ValidatorFunc func(data map[string]interface{}) error
	Option        interface {
		Apply(o *option)
	}

	periodicallyUpdate struct {
		interval time.Duration
	}
	option struct {
		fnValidate         ValidatorFunc
		periodicallyUpdate periodicallyUpdate
	}
)

var defaultOption = option{
	fnValidate: nil,
	periodicallyUpdate: periodicallyUpdate{
		interval: 3 * time.Second,
	},
}

type withValidator ValidatorFunc

func (w withValidator) Apply(o *option) {
	o.fnValidate = ValidatorFunc(w)
}

func WithValidator(fn ValidatorFunc) Option {
	return withValidator(fn)
}
