package main

import (
	"fmt"
	"sync"
)

func main9() {
	var ops int64
	var wg sync.WaitGroup
	var mut sync.Mutex //2 вариянт

	wg.Add(50)
	for i := 0; i < 50; i++ {
		//wg.Add(1) //по группе на каждую горутину
		go func() {
			for i := 0; i < 1000; i++ {
				//atomic.AddInt64(&ops, 1) //чтобы считал правильно

				//2 вариянт
				mut.Lock()
				ops++
				mut.Unlock()
			}
			wg.Done() //
		}()
	}
	wg.Wait() //
	fmt.Println(ops)
}
