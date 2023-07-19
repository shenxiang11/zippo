package slice

import (
	"reflect"
	"testing"
)

func TestFilter_Integers(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	filteredSlice := Filter(intSlice, func(el int, index int) bool {
		return el%2 == 0 // Filter even numbers
	})

	expectedSlice := []int{2, 4}
	if !reflect.DeepEqual(filteredSlice, expectedSlice) {
		t.Errorf("TestFilter_Integers failed. Expected: %v, Got: %v", expectedSlice, filteredSlice)
	}
}

func TestFilter_Strings(t *testing.T) {
	strSlice := []string{"apple", "banana", "orange", "grape"}
	filteredSlice := Filter(strSlice, func(el string, index int) bool {
		return len(el) > 5 // Filter strings with length greater than 5
	})

	expectedSlice := []string{"banana", "orange"}
	if !reflect.DeepEqual(filteredSlice, expectedSlice) {
		t.Errorf("TestFilter_Strings failed. Expected: %v, Got: %v", expectedSlice, filteredSlice)
	}
}

func TestFilter_EmptySlice(t *testing.T) {
	emptySlice := []int{}
	filteredSlice := Filter(emptySlice, func(el int, index int) bool {
		return el > 0 // Filter positive numbers
	})

	// Empty slice remains unchanged
	if len(filteredSlice) != 0 {
		t.Errorf("TestFilter_EmptySlice failed. Expected: empty slice, Got: %v", filteredSlice)
	}
}

func TestFilter_AllElements(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	filteredSlice := Filter(intSlice, func(el int, index int) bool {
		return true // Filter all elements
	})

	// All elements should be present in the filtered slice
	if !reflect.DeepEqual(filteredSlice, intSlice) {
		t.Errorf("TestFilter_AllElements failed. Expected: %v, Got: %v", intSlice, filteredSlice)
	}
}
