package main

import (
	"fmt"
	"time"
)

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func pinger(c chan string) { //if you want to specify the direcation you can add inside the paramer this (c chan <- string)
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) { //if you want to specify the direcation you can add inside the paramer this (c <- chan string)
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	fmt.Println("-----Channels-----")
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
