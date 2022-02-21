package main

import (
	"flag"
	"fmt"

	"github.com/keshiba/ll-dynamicprogramming/pkg/algorithm/fibonacci"
)

func main() {

	var query int64
	flag.Int64Var(&query, "n", 10, "Fibonacci limit")
	flag.Parse()

	result := fibonacci.MemoizedFibonacci(query)
	fmt.Printf("Fib(%v) = %v\n", query, result)

}
