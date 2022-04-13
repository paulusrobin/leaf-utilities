package leafValidatorV10

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
)

type (
	validatorOptions struct {
		translator *ut.UniversalTranslator
	}
	ValidatorOption interface {
		Apply(o *validatorOptions)
	}
)

func defaultValidatorOptions() validatorOptions {
	langEn := en.New()
	langId := id.New()
	return validatorOptions{
		translator: ut.New(langEn, langEn, langId),
	}
}

type withTranslator struct{ *ut.UniversalTranslator }

func (w withTranslator) Apply(o *validatorOptions) {
	o.translator = w.UniversalTranslator
}

func WithTranslator(translator *ut.UniversalTranslator) ValidatorOption {
	return withTranslator{translator}
}
