package main

func Factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}
