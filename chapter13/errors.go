package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("-----Creating a new error-----")
	err := errors.New("Error message")
	fmt.Println(err)
}
