package main

import "fmt"

func PanicFunc() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
}
