package slice

func RemoveAt[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}

	copy(s[index:], s[index+1:])

	result := s[:len(s)-1]

	if cap(result) >= 2*len(result) {
		newResult := make([]T, len(result))
		copy(newResult, result)
		return newResult
	}

	return result
}
