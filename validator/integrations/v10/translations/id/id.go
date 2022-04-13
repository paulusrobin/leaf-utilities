package id

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idDefaultTranslations "github.com/go-playground/validator/v10/translations/id"
	"log"
	"strings"
)

func RegisterTranslations(v *validator.Validate, trans ut.Translator) (err error) {

	if err := idDefaultTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		return err
	}

	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "precision",
			translation: "{0} presisi maksimum adalah {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield_daterange",
			translation: "{0} diluar rentang {1} hari dengan {2}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				tags := strings.Split(fe.Param(), ":")

				t, err := ut.T(fe.Tag(), fe.Field(), tags[1], tags[0])
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield_datetimerange",
			translation: "{0} diluar rentang {1} jam dengan {2}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				tags := strings.Split(fe.Param(), ":")

				t, err := ut.T(fe.Tag(), fe.Field(), tags[1], tags[0])
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)

		} else if t.customTransFunc != nil && t.customRegisFunc == nil {

			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)

		} else if t.customTransFunc == nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)

		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return

	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
