package fibonacci

import (
	"math/big"
)

// BottomUpFibonacci calculates the nth fibonacci number
// using a bottom-up approach as opposed a top-down approach
// which requires recursive function calls to traverse down
// the computation tree. Using a bottom-up approach, we can
// cut down the memory overhead of maintaining a map/cache of
// all values and instead use just 2 variables to store previous
// values
func BottomUpFibonacci(number int64) *big.Int {

	prevValue1 := big.NewInt(0)
	prevValue2 := big.NewInt(1)

	for i := int64(1); i < number; i++ {
		result := new(big.Int).Add(prevValue1, prevValue2)
		prevValue1 = new(big.Int).Set(prevValue2)
		prevValue2 = result
	}

	return prevValue2
}
