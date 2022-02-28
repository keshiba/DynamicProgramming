package changemaking

import (
	"testing"
)

type TestCaseInput struct {
	Denominations []int
	Target        int
}

var testCasesTable = []struct {
	input    TestCaseInput
	expected int
}{
	{
		TestCaseInput{
			[]int{1, 5, 12, 19},
			16,
		},
		4,
	},
	{
		TestCaseInput{
			[]int{1, 5, 16, 25},
			33,
		},
		3,
	},
}

func TestMakeChange(t *testing.T) {

	for _, testCase := range testCasesTable {
		result := MakeChange(
			testCase.input.Denominations, testCase.input.Target)

		if result != testCase.expected {
			t.Errorf(
				"Result %d not equal to expected %d",
				result, testCase.expected)
		}
	}
}
