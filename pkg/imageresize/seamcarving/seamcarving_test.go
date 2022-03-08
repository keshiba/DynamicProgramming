package seamcarving

import (
	"testing"

	"github.com/keshiba/ll-dynamicprogramming/pkg/imageresize/types"
)

var testCaseTable = []struct {
	input    [][]float64
	expected []types.Coordinate
}{
	{
		input: [][]float64{
			{100, 100, 4, 100, 100},
			{100, 10, 5, 2, 100},
			{1, 30, 10, 10, 10},
		},
		expected: []types.Coordinate{
			{Row: 0, Col: 2},
			{Row: 1, Col: 1},
			{Row: 2, Col: 0},
		},
	},
}

func TestComputeVerticalSeam(t *testing.T) {

	for _, test := range testCaseTable {
		coords := ComputeVerticalSeam(test.input)

		if len(coords) != len(test.expected) {
			t.Errorf("Coordinates length %d doesn't match expected length %d",
				len(coords), len(test.expected))
		}

		for i := range coords {

			if coords[i] != test.expected[i] {
				t.Errorf("Coordinate %#v at %d doesn't match expected value of %#v",
					coords[i], i, test.expected[i])
			}
		}
	}
}
