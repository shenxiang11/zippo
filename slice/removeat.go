package slice

import "math"

var defaultSizeBound = 256

// 如果发生缩容，不会缩到这个大小以下，可以根据需要来设置
// 如果这个值相对很小的，那么会比较频繁地触发扩容缩容，影响性能
var sizeBound = defaultSizeBound

func RemoveAt[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}

	copy(s[index:], s[index+1:])

	result := s[:len(s)-1]

	c := cap(result)

	if c >= 2*len(result) && c > sizeBound {
		size := math.Max(float64(len(result)), float64(sizeBound))
		newResult := make([]T, int(size))
		copy(newResult, result)
		return newResult
	}

	return result
}

func SetAdjustBound(x int) {
	sizeBound = x
}

func ResetAdjustBound() {
	sizeBound = defaultSizeBound
}
