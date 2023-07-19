package slice

func Reduce[T any, U any](s []T, cb func(acc U, el T, index int) U, initial U) U {
	value := initial
	for i, el := range s {
		value = cb(value, el, i)
	}
	return value
}
