package fibonacci

import (
	"math/big"
)

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
