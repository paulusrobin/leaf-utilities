package leafValidatorV10

import (
	"database/sql/driver"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	leafTypes "github.com/paulusrobin/leaf-utilities/common/types"
	"github.com/paulusrobin/leaf-utilities/validator/integrations/v10/custom"
	enTranslations "github.com/paulusrobin/leaf-utilities/validator/integrations/v10/translations/en"
	idTranslations "github.com/paulusrobin/leaf-utilities/validator/integrations/v10/translations/id"
	leafValidator "github.com/paulusrobin/leaf-utilities/validator/validator"
	"reflect"
)

type (
	implementation struct {
		instance   *validator.Validate
		translator *ut.UniversalTranslator
	}
)

func New(opts ...ValidatorOption) (leafValidator.Validator, error) {
	options := defaultValidatorOptions()
	for _, opt := range opts {
		opt.Apply(&options)
	}

	enTranslator, _ := options.translator.GetTranslator("en")
	idTranslator, _ := options.translator.GetTranslator("id")

	validate := validator.New()
	if err := enTranslations.RegisterTranslations(validate, enTranslator); err != nil {
		return nil, err
	}
	if err := idTranslations.RegisterTranslations(validate, idTranslator); err != nil {
		return nil, err
	}

	// register all types.Null* types to use the ValidateValuer CustomTypeFunc
	validate.RegisterCustomTypeFunc(ValidateValuer,
		leafTypes.NullString{},
		leafTypes.NullInt32{},
		leafTypes.NullInt64{},
		leafTypes.NullBool{},
		leafTypes.NullFloat64{},
		leafTypes.NullTime{})

	instance := &implementation{instance: validate, translator: options.translator}
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
	if err := i.instance.RegisterValidation("precision", custom.Precision); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("ltecsfield_daterange", custom.IsDateInRangeCrossStructField); err != nil {
		return err
	}
	if err := i.instance.RegisterValidation("ltecsfield_datetimerange", custom.IsDateTimeInRangeCrossStructField); err != nil {
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
