package custom

import (
	"github.com/go-playground/validator/v10"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"strings"
	"time"
)

func IsDateTimeInRangeCrossStructField(fl validator.FieldLevel) bool {
	meta := strings.Split(fl.Param(), ":")
	if len(meta) < 2 {
		return false
	}

	comparedField, _, ok := fl.GetStructFieldOKAdvanced(fl.Parent(), meta[0])
	if !ok {
		return false
	}

	var (
		fieldValueStr            = leafFunctions.ConvertReflectValueToString(fl.Field())
		comparableFieldValueStr  = leafFunctions.ConvertReflectValueToString(comparedField)
		durationValueLimit       = meta[1]
		durationValueLimitNumber = leafFunctions.ConvertStringToUint64(durationValueLimit)
	)

	startDate, err := time.Parse(time.RFC3339, fieldValueStr)
	if err != nil {
		return false
	}

	endDate, err := time.Parse(time.RFC3339, comparableFieldValueStr)
	if err != nil {
		return false
	}

	diff := endDate.Sub(startDate)

	return diff >= 0 && uint64(diff) < durationValueLimitNumber
}
