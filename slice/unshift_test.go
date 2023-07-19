package slice

import (
	"reflect"
	"testing"
)

func TestUnshift_EmptySlice(t *testing.T) {
	emptySlice := []int{}
	emptySlice = Unshift(emptySlice, 1, 2, 3)

	expectedSlice := []int{1, 2, 3}
	if !reflect.DeepEqual(emptySlice, expectedSlice) {
		t.Errorf("TestUnshift_EmptySlice failed. Expected: %v, Got: %v", expectedSlice, emptySlice)
	}
}

func TestUnshift_NonEmptySlice(t *testing.T) {
	intSlice := []int{2, 3, 4}
	intSlice = Unshift(intSlice, 1)

	expectedSlice := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(intSlice, expectedSlice) {
		t.Errorf("TestUnshift_NonEmptySlice failed. Expected: %v, Got: %v", expectedSlice, intSlice)
	}
}

func TestUnshift_MultipleElements(t *testing.T) {
	strSlice := []string{"b", "c"}
	strSlice = Unshift(strSlice, "a", "d", "e")

	expectedSlice := []string{"a", "d", "e", "b", "c"}
	if !reflect.DeepEqual(strSlice, expectedSlice) {
		t.Errorf("TestUnshift_MultipleElements failed. Expected: %v, Got: %v", expectedSlice, strSlice)
	}
}
