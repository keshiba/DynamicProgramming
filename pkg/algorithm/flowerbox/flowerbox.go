package flowerbox

import "fmt"

// CalculateMaxFlowerBoxHeight computes the maximum height
// of flowers in a box with certain constraints according
// to the FlowerBox problem
// 1. Flowers should not be planted adjacently
//		If there's a flower at i, then the next flower
//		should be at i-2
// 2. The height of a flower is equal to the nutrients
//		at ith location it was planted
// Goal: Maximize the height of flowers by choosing
//		optimal locations to plant them
// Problem is similar to bottomup fibonacci and can
// be simplified to the following recursive function
// f(i) = max( nutrient_at_i + f (i-2), f(i-1))
// f(i) -> max height of flowers at location i
// At any location i, we can either choose to plant at that
// location, or to skip that location and move to the next one
// If we plant at a location, then we sum the nutrients at i and
// skip 2 places according to the problem's rules (no adjacent flowers)
// Else, if we move to the next location, then we calculate max
// from that location
func CalculateMaxFlowerBoxHeight(nutrients []int) int {

	prevMaxNutrient1 := 0 // f (i-1)
	prevMaxNutrient2 := 0 // f (i-2)

	for _, nutrient := range nutrients {
		maxNutrient := max(nutrient+prevMaxNutrient2, prevMaxNutrient1)
		prevMaxNutrient2 = prevMaxNutrient1
		prevMaxNutrient1 = maxNutrient

		fmt.Printf(
			"N=%d, M=%d, p1=%d, p2=%d\n",
			nutrient, maxNutrient, prevMaxNutrient1, prevMaxNutrient2)
	}

	return prevMaxNutrient1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
