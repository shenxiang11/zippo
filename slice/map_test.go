package slice

import (
	"reflect"
	"testing"
)

type A struct {
	val int64
}

type B struct {
	value int64
}

func TestMap(t *testing.T) {
	arrA := []A{
		{val: 10},
		{val: 20},
	}

	result := Map[A, B](arrA, func(el A, index int) B {
		return B{value: el.val * 2}
	})

	expectedSlice := []B{
		{value: 20},
		{value: 40},
	}
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("TestMap failed. Expected: %v, Got: %v", expectedSlice, result)
	}
}
