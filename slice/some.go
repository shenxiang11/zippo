package slice

func Some[T any](s []T, check func(el T, index int) bool) bool {
	for i, el := range s {
		tmp := check(el, i)
		if tmp {
			return true
		}
	}

	return false
}
