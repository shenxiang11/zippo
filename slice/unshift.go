package slice

func Unshift[T any](s []T, el ...T) []T {
	newSlice := make([]T, len(s)+len(el))
	copy(newSlice, el)
	copy(newSlice[len(el):], s)
	return newSlice
}
