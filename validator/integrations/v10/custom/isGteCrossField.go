package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"strconv"
)

func IsGteCrossStructField(fl validator.FieldLevel) bool {
	topField, _, ok := fl.GetStructFieldOK()
	if !ok {
		return false
	}

	fieldValue, err := strconv.Atoi(leafFunctions.ConvertReflectValueToString(fl.Field()))
	if err != nil {
		return false
	}

	topValue, err := strconv.Atoi(leafFunctions.ConvertReflectValueToString(topField))
	if err != nil {
		return false
	}

	return fieldValue >= topValue
}
