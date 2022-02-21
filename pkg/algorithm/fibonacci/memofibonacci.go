package fibonacci

import (
	"math/big"
)

var memo map[string]*big.Int

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
