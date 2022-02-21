package fibonacci

import (
	"fmt"
	"math/big"
	"testing"
)

var smallInputTable = []struct {
	input    uint64
	expected uint64
}{
	{input: 5, expected: 5},
	{input: 10, expected: 55},
	{input: 20, expected: 6765},
	{input: 30, expected: 832040},
	{input: 40, expected: 102334155},
}

var largeInputTable = []struct {
	input    int64
	expected func() *big.Int
}{
	{input: 5, expected: func() *big.Int { return big.NewInt(5) }},
	{input: 10, expected: func() *big.Int { return big.NewInt(55) }},
	{input: 20, expected: func() *big.Int { return big.NewInt(6765) }},
	{input: 30, expected: func() *big.Int { return big.NewInt(832040) }},
	{input: 40, expected: func() *big.Int { return big.NewInt(102334155) }},
	{input: 50, expected: func() *big.Int { return big.NewInt(12586269025) }},
	{input: 70, expected: func() *big.Int { return big.NewInt(117669030460994 + 72723460248141) }},
	{input: 99, expected: func() *big.Int {
		val := new(big.Int)
		val.SetString("218922995834555169026", 10)
		return val
	}},
	{input: 100, expected: func() *big.Int {
		val := new(big.Int)
		val.SetString("354224848179261915075", 10)
		return val
	}},
	{input: 150, expected: func() *big.Int {
		val := new(big.Int)
		val.SetString("9969216677189303386214405760200", 10)
		return val
	}},
	{input: 200, expected: func() *big.Int {
		val := new(big.Int)
		val.SetString("280571172992510140037611932413038677189525", 10)
		return val
	}},
	{input: 300, expected: func() *big.Int {
		val := new(big.Int)
		val.SetString("222232244629420445529739893461909967206666939096499764990979600", 10)
		return val
	}},
}

func TestFibonacci(t *testing.T) {

	for _, test := range smallInputTable {
		result := Fibonacci(test.input)

		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestMemoizedFibonacci(t *testing.T) {

	for _, test := range largeInputTable {
		result := MemoizedFibonacci(test.input)

		if result.Cmp(test.expected()) != 0 {
			t.Errorf("Result %s not equal to expected %s", result.Text(10), test.expected().Text(10))
		}
	}
}

func TestBottomUpFibonacci(t *testing.T) {

	for _, test := range largeInputTable {
		result := BottomUpFibonacci(test.input)

		if result.Cmp(test.expected()) != 0 {
			t.Errorf("Result %s not equal to expected %s", result.Text(10), test.expected().Text(10))
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {

	for _, v := range smallInputTable {

		tcName := fmt.Sprintf("input_%v", v.input)

		b.Run(tcName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fibonacci(v.input)
			}
		})
	}

}

func BenchmarkMemoizedFibonacci(b *testing.B) {

	for _, v := range largeInputTable {

		tcName := fmt.Sprintf("input_%v", v.input)

		b.Run(tcName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MemoizedFibonacci(v.input)
			}
		})
	}
}

func BenchmarkBottomUpFibonacci(b *testing.B) {

	for _, v := range largeInputTable {

		tcName := fmt.Sprintf("input_%v", v.input)

		b.Run(tcName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BottomUpFibonacci(v.input)
			}
		})
	}
}
