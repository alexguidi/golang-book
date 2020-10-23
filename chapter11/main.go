package main

import (
	"fmt"
	"golang-book/chapter11/math" //if we want to use two packages with the same name (like math) we can define a alias like import m "golang-book/chapter11/math" where m is a alias
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
