package slice

func Unique[T int | int32 | int64 | string | float32 | float64](s []T) []T {
	var result []T
	temp := map[T]byte{}
	for _, e := range s {
		l := len(temp)
		temp[e] = 0
		if len(temp) != l {
			result = append(result, e)
		}
	}
	return result
}

func InArray[T int | int32 | int64 | string | float32 | float64](s T, t []T) bool {
	for _, item := range t {
		if s == item {
			return true
		}
	}

	return false
}
