package leafFunctions

func IndexString(data []string, target string) int {
	for i, j := 0, len(data)-1; i <= j; i, j = i+1, j-1 {
		if data[i] == target {
			return i
		}
		if data[j] == target {
			return j
		}
	}
	return -1
}

func IndexInt(data []int, target int) int {
	for i, j := 0, len(data)-1; i <= j; i, j = i+1, j-1 {
		if data[i] == target {
			return i
		}
		if data[j] == target {
			return j
		}
	}
	return -1
}

func IndexInt64(data []int64, target int64) int {
	for i, j := 0, len(data)-1; i <= j; i, j = i+1, j-1 {
		if data[i] == target {
			return i
		}
		if data[j] == target {
			return j
		}
	}
	return -1
}

func IndexUint64(data []uint64, target uint64) int {
	for i, j := 0, len(data)-1; i <= j; i, j = i+1, j-1 {
		if data[i] == target {
			return i
		}
		if data[j] == target {
			return j
		}
	}
	return -1
}

func IndexFloat64(data []float64, target float64) int {
	for i, j := 0, len(data)-1; i <= j; i, j = i+1, j-1 {
		if data[i] == target {
			return i
		}
		if data[j] == target {
			return j
		}
	}
	return -1
}
