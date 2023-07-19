package slice

import (
	"reflect"
	"testing"
)

func TestToReversed_Integers(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	reversedSlice := ToReversed(intSlice)

	expectedSlice := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(reversedSlice, expectedSlice) {
		t.Errorf("TestToReversed_Integers failed. Expected: %v, Got: %v", expectedSlice, reversedSlice)
	}

	// Ensure the original slice remains unchanged
	if !reflect.DeepEqual(intSlice, []int{1, 2, 3, 4, 5}) {
		t.Error("TestToReversed_Integers failed. The original slice was modified.")
	}
}

func TestToReversed_Strings(t *testing.T) {
	strSlice := []string{"apple", "banana", "orange", "grape"}
	reversedSlice := ToReversed(strSlice)

	expectedSlice := []string{"grape", "orange", "banana", "apple"}
	if !reflect.DeepEqual(reversedSlice, expectedSlice) {
		t.Errorf("TestToReversed_Strings failed. Expected: %v, Got: %v", expectedSlice, reversedSlice)
	}

	// Ensure the original slice remains unchanged
	if !reflect.DeepEqual(strSlice, []string{"apple", "banana", "orange", "grape"}) {
		t.Error("TestToReversed_Strings failed. The original slice was modified.")
	}
}

func TestToReversed_EmptySlice(t *testing.T) {
	emptySlice := []int{}
	reversedSlice := ToReversed(emptySlice)

	// Empty slice remains unchanged when reversed
	if !reflect.DeepEqual(reversedSlice, emptySlice) {
		t.Errorf("TestToReversed_EmptySlice failed. Expected: %v, Got: %v", emptySlice, reversedSlice)
	}
}

func TestToReversed_SingleElementSlice(t *testing.T) {
	singleElementSlice := []int{42}
	reversedSlice := ToReversed(singleElementSlice)

	// Single-element slice remains unchanged when reversed
	if !reflect.DeepEqual(reversedSlice, singleElementSlice) {
		t.Errorf("TestToReversed_SingleElementSlice failed. Expected: %v, Got: %v", singleElementSlice, reversedSlice)
	}
}
