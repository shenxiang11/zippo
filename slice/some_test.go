package slice

import "testing"

func TestSome_Integers(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	checkFunc := func(el int, index int) bool {
		return el%2 == 0 // Check if the element is even
	}
	result := Some(intSlice, checkFunc)
	if !result {
		t.Errorf("TestSome_Integers failed. Expected: true, Got: false")
	}
}

func TestSome_Strings(t *testing.T) {
	strSlice := []string{"apple", "banana", "orange", "grape"}
	checkFunc := func(el string, index int) bool {
		return el == "orange" // Check if the element is "orange"
	}
	result := Some(strSlice, checkFunc)
	if !result {
		t.Errorf("TestSome_Strings failed. Expected: true, Got: false")
	}
}

func TestSome_EmptySlice(t *testing.T) {
	emptySlice := []int{}
	checkFunc := func(el int, index int) bool {
		return el == 0 // Check if the element is 0
	}
	result := Some(emptySlice, checkFunc)
	if result {
		t.Errorf("TestSome_EmptySlice failed. Expected: false, Got: true")
	}
}
