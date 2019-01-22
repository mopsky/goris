package utils

func AbsInt(value int) int {
	if value >= 0 {
		return value
	} else {
		return -value
	}
}

func AbsInt64(value int64) int64 {
	if value >= 0 {
		return value
	} else {
		return -value
	}
}