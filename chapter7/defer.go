package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func DeferFunc() {
	defer second() //defer schedules a function to be call after the function completes, in this case second() will be called when the DeferFunc finishes
	first()
}
