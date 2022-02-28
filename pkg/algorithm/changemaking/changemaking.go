package changemaking

import (
	"math"
)

type Pair struct {
	index  int
	target int
}

var cache map[Pair]int
var denominations []int

func MakeChange(denominationsArg []int, target int) int {
	denominations = denominationsArg
	cache = make(map[Pair]int)

	return makeChangeSubProblem(len(denominations)-1, target)
}

func makeChangeSubProblem(changeIndex, target int) int {

	// f(i, t) = min ( f(i, t-di) + 1,  f(i-1, t) )
	var resultIfDenominationUsed, resultIfDenominationIgnored int
	key := Pair{index: changeIndex, target: target}

	if result, exists := cache[key]; exists {
		return result
	}

	currentDenomination := denominations[changeIndex]

	if currentDenomination > target {
		resultIfDenominationUsed = math.MaxInt
	} else if currentDenomination == target {
		resultIfDenominationUsed = 1
	} else {
		newTarget := target - currentDenomination
		resultIfDenominationUsed = 1 + makeChangeSubProblem(
			changeIndex, newTarget)
	}

	if changeIndex == 0 {
		resultIfDenominationIgnored = math.MaxInt
	} else {

		resultIfDenominationIgnored = makeChangeSubProblem(
			changeIndex-1, target)
	}

	result := min(resultIfDenominationUsed, resultIfDenominationIgnored)
	cache[key] = result

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
