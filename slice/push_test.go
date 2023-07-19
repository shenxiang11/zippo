package slice

import (
	"reflect"
	"testing"
)

func TestPush_EmptySlice(t *testing.T) {
	intSlice := []int{}
	intSlice = Push(intSlice, 1, 2, 3)
	expectedIntSlice := []int{1, 2, 3}
	if !reflect.DeepEqual(intSlice, expectedIntSlice) {
		t.Errorf("Test case failed. Expected: %v, Got: %v", expectedIntSlice, intSlice)
	}
}

func TestPush_NonEmptySlice(t *testing.T) {
	strSlice := []string{"a", "b"}
	strSlice = Push(strSlice, "c", "d", "e")
	expectedStrSlice := []string{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(strSlice, expectedStrSlice) {
		t.Errorf("Test case failed. Expected: %v, Got: %v", expectedStrSlice, strSlice)
	}
}

func TestPush_NilSlice(t *testing.T) {
	var floatSlice []float64
	floatSlice = Push(floatSlice, 1.1, 2.2, 3.3)
	expectedFloatSlice := []float64{1.1, 2.2, 3.3}
	if !reflect.DeepEqual(floatSlice, expectedFloatSlice) {
		t.Errorf("Test case failed. Expected: %v, Got: %v", expectedFloatSlice, floatSlice)
	}
}
