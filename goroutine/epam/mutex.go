package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main6() {
	ch := make(chan int)
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("END")
}

func work(n int, ch chan int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	//counter = 0
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("mutex N: ", n, " counter: ", counter)
	}

	ch <- 0
}
