package main

import "fmt"

// this is a comment

var (
	a1 = 5
	b1 = 10
	c1 = 15
	d1 = "asasd"
)

const (
	a2 = 5
	b2 = 10
	c2 = 15
	d2 = "aaaa"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hi Alex"
	fmt.Println(y)

	const z string = "Hi Nyc"
	fmt.Println(z)

	fmt.Println(d1)
	fmt.Println(d2)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

}
