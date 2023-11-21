package utils

func Filter[T any](sl []T, pred func(el *T) bool) []T {
	res := make([]T, 0, len(sl))
	for ind := 0; ind < len(sl); ind++ {
		if pred(&sl[ind]) {
			res = append(res, sl[ind])
		}
	}
	return res
}

func Map[T any, U any](sl []T, trans func(el *T) U) []U {
	res := make([]U, 0, len(sl))
	for ind := 0; ind < len(sl); ind++ {
		res = append(res, trans(&sl[ind]))
	}
	return res
}

func MaxStringLen[T any](sl []T, get func(el *T) string) int {
	var max int
	for ind := 0; ind < len(sl); ind++ {
		leng := len(get(&sl[ind]))
		if leng > max {
			max = leng
		}
	}
	return max
}
