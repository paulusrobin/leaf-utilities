package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"time"
)

const dateLayout = "2006-01-02"

func IsDate(fl validator.FieldLevel) bool {
	checkedValue := leafFunctions.ConvertReflectValueToString(fl.Field())

	_, err := time.Parse(dateLayout, checkedValue)
	if err != nil {
		return false
	}

	return true
}
