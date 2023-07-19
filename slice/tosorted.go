package slice

import "sort"

func ToSorted[T any](s []T, compare func(el1, el2 T) bool) []T {
	sortedSlice := make([]T, len(s))
	copy(sortedSlice, s)

	sort.Slice(sortedSlice, func(i, j int) bool {
		return compare(sortedSlice[i], sortedSlice[j])
	})

	return sortedSlice
}
