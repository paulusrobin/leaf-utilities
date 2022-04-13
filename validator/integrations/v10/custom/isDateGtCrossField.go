package custom

import (
	"github.com/araddon/dateparse"
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
)

func IsDateGtCrossStructField(fl validator.FieldLevel) bool {
	topField, _, ok := fl.GetStructFieldOK()
	if !ok {
		return false
	}

	fieldTime, err := dateparse.ParseAny(leafFunctions.ConvertReflectValueToString(fl.Field()))
	if err != nil {
		return false
	}

	topTime, err := dateparse.ParseAny(leafFunctions.ConvertReflectValueToString(topField))
	if err != nil {
		return false
	}

	return fieldTime.After(topTime) || fieldTime.Equal(topTime)
}
