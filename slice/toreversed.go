package slice

func ToReversed[T any](s []T) []T {
	l := len(s)
	result := make([]T, l)
	for i, item := range s {
		result[l-1-i] = item
	}

	return result
}
