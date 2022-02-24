package leafFunctions

import "math"

func CalculateTotalPages(dataTotalCount int, limit int) int {
	totalPagesInFloat := float64(dataTotalCount) / float64(limit)
	totalPagesInFloat = math.Ceil(totalPagesInFloat)
	return int(totalPagesInFloat)
}
