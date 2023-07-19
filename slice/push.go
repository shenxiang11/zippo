package slice

func Push[T any](s []T, el ...T) []T {
	return append(s, el...)
}
