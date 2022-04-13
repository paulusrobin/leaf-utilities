package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"reflect"
	"strings"
)

func IsRequiredIf(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), ";")
	paramsLength := len(params)
	value := leafFunctions.ConvertReflectValueToString(fl.Field())

	if paramsLength >= 3 {
		checkedColumnValue := leafFunctions.ConvertReflectValueToString(reflect.Indirect(fl.Parent()).FieldByName(params[0]))
		expression := params[1]
		var checkedValue string
		isRequired := false

		if params[2][0] == '\'' {
			checkedValue = strings.TrimLeft(params[2], "'")
		} else {
			checkedValue = leafFunctions.ConvertReflectValueToString(reflect.Indirect(fl.Parent()).FieldByName(params[2]))
		}

		if expression == "=" {
			isRequired = checkedColumnValue == checkedValue
		} else {
			isRequired = checkedColumnValue != checkedValue
		}

		if isRequired {
			return value != "" && value != "0"
		}

		return true
	}

	return false
}
