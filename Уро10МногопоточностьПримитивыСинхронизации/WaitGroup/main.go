package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter uint64
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ //поконебезопастность поэтому тут не 1000 а разные числа
				// atomic. Addlint64(&counter, 1) //а вот атомик потокобезопасен
			}
		}()
	}
	wg.Wait()
	fmt.Printf("all done, counter=%d\n", counter)

}

/*
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) //говорим что будет запущена одна горутина
		k := i
		go func() {
			defer wg.Done() //делаем ее
			fmt.Printf("%d goroutine working…\n", k)
			time.Sleep(300 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Println(" all done ")

}
*/
