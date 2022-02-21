package fibonacci

// Fibonacci calculates the nth fibonacci number using
// a traditional recursive algorithm.
// This is to demonstrate the inefficiencies of an unoptimized
// recursive algorithm
func Fibonacci(number uint64) uint64 {

	if number <= 1 {
		return number
	} else {
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}
