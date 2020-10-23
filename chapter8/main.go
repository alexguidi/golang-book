package main

import "fmt"

func zeroByValue(x int) {
	x = 0
}

func zeroByPointer(xPtr *int) /*pointer to int*/ {
	*xPtr = 0 //this means store the 0 value in the memory location xPtr refers to
}

func main() {
	x := 5

	zeroByValue(x) //passing by value
	fmt.Println(x) // x is still 5

	zeroByPointer(&x) //passing by reference
	fmt.Println(x)

	xPtr := new(int) //another way to create a pointer
	zeroByPointer(xPtr)
	fmt.Println(*xPtr)

}
