package leafValidatorV10

import (
	"database/sql/driver"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	leafTypes "github.com/paulusrobin/leaf-utilities/common/types"
	"github.com/paulusrobin/leaf-utilities/validator/integrations/v10/custom"
	leafValidator "github.com/paulusrobin/leaf-utilities/validator/validator"
	"reflect"
)

type (
	implementation struct {
		instance *validator.Validate
		trans    ut.Translator
	}
)

func New() (leafValidator.Validator, error) {
	langEn := en.New()
	langId := id.New()
	uni := ut.New(langEn, langEn, langId)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}

	// register all types.Null* types to use the ValidateValuer CustomTypeFunc
	validate.RegisterCustomTypeFunc(ValidateValuer, leafTypes.NullString{}, leafTypes.NullInt32{}, leafTypes.NullInt64{}, leafTypes.NullBool{}, leafTypes.NullFloat64{}, leafTypes.NullTime{})

	instance := &implementation{instance: validate, trans: trans}
	if err := instance.registerDefaultValidator(); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *implementation) registerDefaultValidator() error {
	if err := i.instance.RegisterValidation("date", custom.IsDate); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("datetime", custom.IsDateTime); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("date_gtecsfield", custom.IsDateGtCrossStructField); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("max_if", custom.IsMaxIf); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("required_if", custom.IsRequiredIf); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("gtecsfield", custom.IsGteCrossStructField); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("gtcsfield", custom.IsGtCrossStructField); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("ltecsfield", custom.IsLteCrossStructField); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("ltcsfield", custom.IsLtCrossStructField); err != nil {
		return err
	}
	return nil
}

func (i *implementation) RegisterValidation(tag string, fn func(fl validator.FieldLevel) bool) {
	i.instance.RegisterValidation(tag, fn)
}

func (i *implementation) RegisterStructValidation(fn func(sl validator.StructLevel), types interface{}) {
	i.instance.RegisterStructValidation(fn, types)
}

func (i *implementation) Validate(object interface{}) error {
	if err := i.instance.Struct(object); err != nil {
		return err
	}
	return nil
}

func (i *implementation) ValidateVar(object interface{}, constraint string) error {
	if err := i.instance.Var(object, constraint); err != nil {
		return err
	}
	return nil
}

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}
	return nil
}
