package main

import (
	"fmt"
	"strconv"
)

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func f2() (r int) {
	r = 1
	return
}

func f() (int, int, int) {
	return 5, 6, 7
}

func add(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}

func main() {
	fl := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(fl))
	fmt.Println(f2())

	x, y, z := f()
	fmt.Println(x, y, z)

	fmt.Println(add(1, 1, 1, 1, 2))
	xs := []int{1, 2, 3, 2}
	fmt.Println(add(xs...))

	fmt.Println("-----Closure-----")
	Closure()

	fmt.Println("-----Recursion-----")
	var uval uint64 = Factorial(2)
	fmt.Println(strconv.FormatUint(uval, 10))

	fmt.Println("-----Defer-----")
	DeferFunc()

	fmt.Println("-----Panic and Recover-----")
	PanicFunc()
}
