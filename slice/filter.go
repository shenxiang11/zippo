package slice

func Filter[T any](s []T, f func(el T, index int) bool) []T {
	result := []T{}

	for i, el := range s {
		t := f(el, i)
		if t {
			result = append(result, el)
		}
	}

	return result
}
