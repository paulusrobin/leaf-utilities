package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"strconv"
	"strings"
)

func Precision(fl validator.FieldLevel) bool {
	precision, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	fieldValue := leafFunctions.ConvertReflectValueToString(fl.Field())

	splits := strings.Split(fieldValue, ".")
	return len(splits) == 1 || len(splits) > 1 && len(splits[1]) <= precision
}
