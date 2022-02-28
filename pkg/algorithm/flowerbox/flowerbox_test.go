package flowerbox

import (
	"testing"
)

var testcaseTable = []struct {
	input    []int
	expected int
}{
	{input: []int{3, 10, 3, 1, 2}, expected: 12},
	{input: []int{9, 10, 9}, expected: 18},
}

func TestCalculateMaxFlowerBoxHeight(t *testing.T) {

	for _, testcase := range testcaseTable {
		result := CalculateMaxFlowerBoxHeight(testcase.input)

		if result != testcase.expected {
			t.Errorf(
				"Result %d not equal to expected %d",
				result, testcase.expected)

		}
	}
}
