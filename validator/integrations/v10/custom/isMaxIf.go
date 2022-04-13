package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"reflect"
	"strings"
)

func IsMaxIf(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), ";")
	paramsLength := len(params)
	value := leafFunctions.ConvertReflectValueToString(fl.Field())

	if paramsLength >= 4 {
		maxValue := params[0]
		checkedColumnValue := leafFunctions.ConvertReflectValueToString(reflect.Indirect(fl.Parent()).FieldByName(params[1]))
		expression := params[2]
		var checkedValue string
		isRequired := false

		if params[3][0] == '\'' {
			checkedValue = strings.TrimLeft(params[3], "'")
		} else {
			checkedValue = leafFunctions.ConvertReflectValueToString(reflect.Indirect(fl.Parent()).FieldByName(params[3]))
		}

		if expression == "=" {
			isRequired = checkedColumnValue == checkedValue
		} else {
			isRequired = checkedColumnValue != checkedValue
		}

		if isRequired {
			return leafFunctions.ConvertStringToUint64(value) <= leafFunctions.ConvertStringToUint64(maxValue)
		}

		return true
	}

	return false
}
