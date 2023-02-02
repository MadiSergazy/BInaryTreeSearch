package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func writeChan(c0, c1 chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c0)
	}
	c1 <- 0
}

func main1() {
	c := make(chan int)
	quit := make(chan int)
	go writeChan(c, quit)
	fibonacci(c, quit)
}
