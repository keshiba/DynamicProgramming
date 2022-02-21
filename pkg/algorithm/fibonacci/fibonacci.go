package fibonacci

func Fibonacci(number uint64) uint64 {

	if number <= 1 {
		return number
	} else {
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}
