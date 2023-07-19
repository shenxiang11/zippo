package slice

import "testing"

func TestReduce_IntegerAddition(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	intSum := Reduce(intSlice, func(acc, el int, index int) int {
		return acc + el
	}, 0)

	expectedIntSum := 15
	if intSum != expectedIntSum {
		t.Errorf("Integer addition test failed. Expected: %d, Got: %d", expectedIntSum, intSum)
	}
}

func TestReduce_StringConcatenation(t *testing.T) {
	strSlice := []string{"Hello", " ", "World", "!"}
	strConcat := Reduce(strSlice, func(acc, el string, index int) string {
		return acc + el
	}, "")

	expectedStrConcat := "Hello World!"
	if strConcat != expectedStrConcat {
		t.Errorf("String concatenation test failed. Expected: %s, Got: %s", expectedStrConcat, strConcat)
	}
}

func TestReduce_FloatingPointSum(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	floatSum := Reduce(floatSlice, func(acc, el float64, index int) float64 {
		return acc + el
	}, 0.0)

	expectedFloatSum := 16.5
	if floatSum != expectedFloatSum {
		t.Errorf("Floating-point sum test failed. Expected: %f, Got: %f", expectedFloatSum, floatSum)
	}
}
