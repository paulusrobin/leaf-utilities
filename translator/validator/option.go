package leafTranslatorValidator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
)

type (
	options struct {
		defaultLanguage   locales.Translator
		supportedLanguage []locales.Translator
	}
	Option interface {
		Apply(opt *options)
	}
)

func defaultOption() options {
	localeEn := en.New()
	localeId := id.New()
	return options{
		defaultLanguage:   localeEn,
		supportedLanguage: []locales.Translator{localeEn, localeId},
	}
}
