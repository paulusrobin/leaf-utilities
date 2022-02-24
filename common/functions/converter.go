package leafFunctions

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertReflectValueToString(value reflect.Value) string {
	stringValue := ""

	switch value.Kind() {
	case reflect.String:
		stringValue = value.Interface().(string)
	case reflect.Bool:
		stringValue = strconv.FormatBool(value.Interface().(bool))
	case reflect.Uint64:
		stringValue = strconv.FormatUint(value.Interface().(uint64), 10)
	default:
		stringValue = fmt.Sprint(value)
	}

	return stringValue
}

func ConvertStringToUint64(value string, defaultVal ...uint64) uint64 {
	var defaultValue uint64 = 0
	if len(defaultVal) > 0 {
		defaultValue = defaultVal[0]
	}

	if val, err := strconv.ParseUint(value, 10, 64); err != nil {
		return defaultValue
	} else {
		return val
	}
}

func ConvertUint64ToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}
