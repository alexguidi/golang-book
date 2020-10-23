package main

import "fmt"

func Closure() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(9, 9))
}
