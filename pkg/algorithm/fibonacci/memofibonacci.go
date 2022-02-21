package fibonacci

import (
	"math/big"
)

var memo map[string]*big.Int

// MemoizedFibonacci calculates the nth fibonacci number using
// a map as a value-cache to store computed data for re-use.
// This is generally an optimization over the traditional
// recursive fibonacci algorithm, but it trades-off memory
// in favor of better performance
func MemoizedFibonacci(number int64) big.Int {

	memo = make(map[string]*big.Int)
	bigNum := big.NewInt(number)

	result := MemoizedFibonacciInternal(bigNum)

	return *result
}

func MemoizedFibonacciInternal(number *big.Int) *big.Int {

	if number.Cmp(big.NewInt(1)) <= 0 {

		return number

	} else if result, ok := memo[number.String()]; ok {

		return result

	} else {

		sub1Value := new(big.Int).Sub(number, big.NewInt(1))
		sub2Value := new(big.Int).Sub(number, big.NewInt(2))

		result := big.NewInt(0).Add(
			MemoizedFibonacciInternal(sub1Value),
			MemoizedFibonacciInternal(sub2Value),
		)

		memo[number.String()] = result

		return result
	}
}
