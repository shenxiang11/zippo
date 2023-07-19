package slice

import (
	"reflect"
	"testing"
)

func TestToSorted_Integers(t *testing.T) {
	intSlice := []int{3, 1, 4, 2, 5}
	sortedSlice := ToSorted(intSlice, func(i, j int) bool {
		return i < j
	})

	expectedSlice := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(sortedSlice, expectedSlice) {
		t.Errorf("TestToSorted_Integers failed. Expected: %v, Got: %v", expectedSlice, sortedSlice)
	}
}

func TestToSorted_Strings(t *testing.T) {
	strSlice := []string{"apple", "banana", "orange", "grape"}
	sortedSlice := ToSorted(strSlice, func(i, j string) bool {
		return i < j
	})

	expectedSlice := []string{"apple", "banana", "grape", "orange"}
	if !reflect.DeepEqual(sortedSlice, expectedSlice) {
		t.Errorf("TestToSorted_Strings failed. Expected: %v, Got: %v", expectedSlice, sortedSlice)
	}
}

func TestToSorted_CustomSort(t *testing.T) {
	// Custom sorting for a struct with a 'value' field.
	type Item struct {
		value int
	}
	itemSlice := []Item{
		{value: 3},
		{value: 1},
		{value: 4},
		{value: 2},
		{value: 5},
	}
	sortedSlice := ToSorted(itemSlice, func(i, j Item) bool {
		return i.value > j.value
	})

	expectedSlice := []Item{
		{value: 5},
		{value: 4},
		{value: 3},
		{value: 2},
		{value: 1},
	}
	if !reflect.DeepEqual(sortedSlice, expectedSlice) {
		t.Errorf("TestToSorted_CustomSort failed. Expected: %v, Got: %v", expectedSlice, sortedSlice)
	}
}
