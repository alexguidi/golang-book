package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) //channels are synchronous, both sides of the channel will wait until the other side is ready
	c2 := make(chan string) //but if you want a asynchronous behavior you can create a buffered channel passing a one more parameter, the number one like this c2 := make(chan string, 1)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				/*default:
				fmt.Println("nothing ready")*/
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}
