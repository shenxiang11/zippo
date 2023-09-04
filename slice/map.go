package slice

func Map[T any, U any](s []T, f func(el T, index int) U) []U {
	result := []U{}

	for i, el := range s {
		t := f(el, i)

		result = append(result, t)
	}

	return result
}
