package slice

import (
	"reflect"
	"testing"
)

func TestReverse_Integers(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	Reverse(intSlice)

	expectedSlice := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(intSlice, expectedSlice) {
		t.Errorf("TestReverse_Integers failed. Expected: %v, Got: %v", expectedSlice, intSlice)
	}
}

func TestReverse_Strings(t *testing.T) {
	strSlice := []string{"apple", "banana", "orange", "grape"}
	Reverse(strSlice)

	expectedSlice := []string{"grape", "orange", "banana", "apple"}
	if !reflect.DeepEqual(strSlice, expectedSlice) {
		t.Errorf("TestReverse_Strings failed. Expected: %v, Got: %v", expectedSlice, strSlice)
	}
}

func TestReverse_EmptySlice(t *testing.T) {
	emptySlice := []int{}
	Reverse(emptySlice)

	// Empty slice remains unchanged
	if len(emptySlice) != 0 {
		t.Errorf("TestReverse_EmptySlice failed. Expected: empty slice, Got: %v", emptySlice)
	}
}

func TestReverse_SingleElementSlice(t *testing.T) {
	singleElementSlice := []int{42}
	Reverse(singleElementSlice)

	// Single-element slice remains unchanged
	if singleElementSlice[0] != 42 {
		t.Errorf("TestReverse_SingleElementSlice failed. Expected: [42], Got: %v", singleElementSlice)
	}
}
