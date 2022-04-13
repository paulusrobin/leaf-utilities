package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"time"
)

const (
	dateTimeLayout    = "2006-01-02T15:04:05Z"
	dateTimeAltLayout = "2006-01-02 15:04:05"
)

func IsDateTime(fl validator.FieldLevel) bool {
	checkedValue := leafFunctions.ConvertReflectValueToString(fl.Field())

	_, err := time.Parse(dateTimeLayout, checkedValue)
	if err == nil {
		return true
	}

	_, err = time.Parse(dateTimeAltLayout, checkedValue)
	if err == nil {
		return true
	}

	return false
}
