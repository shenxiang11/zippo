package slice

import (
	"reflect"
	"testing"
)

func TestRemoveAt_IntegerSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	expectedIntSlice := []int{1, 2, 4, 5}
	intSlice = RemoveAt(intSlice, 2)
	if !reflect.DeepEqual(intSlice, expectedIntSlice) {
		t.Errorf("Test case 1 failed. Expected: %v, Got: %v", expectedIntSlice, intSlice)
	}
}

func TestRemoveAt_StringSlice(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d"}
	expectedStrSlice := []string{"b", "c", "d"}
	strSlice = RemoveAt(strSlice, 0)
	if !reflect.DeepEqual(strSlice, expectedStrSlice) {
		t.Errorf("Test case 2 failed. Expected: %v, Got: %v", expectedStrSlice, strSlice)
	}
}

func TestRemoveAt_FloatSlice(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.3}
	floatSlice = RemoveAt(floatSlice, len(floatSlice)-1)
	if !reflect.DeepEqual(floatSlice, expectedFloatSlice) {
		t.Errorf("Test case 3 failed. Expected: %v, Got: %v", expectedFloatSlice, floatSlice)
	}
}

func TestRemoveAt_EmptySlice(t *testing.T) {
	emptySlice := []int{}
	emptySlice = RemoveAt(emptySlice, 0)
	if len(emptySlice) != 0 {
		t.Errorf("Test case 4 failed. Expected: Empty slice, Got: %v", emptySlice)
	}
}

func TestRemoveAt_CapChange(t *testing.T) {
	SetAdjustBound(8)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		slice = RemoveAt(slice, 0)
	}
	if cap(slice) != 8 {
		t.Errorf("Test case 5 failed. Expected: cap(%d), Got: cap(%d)", 5, cap(slice))
	}
	ResetAdjustBound()
}
