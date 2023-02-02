package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main8() {
	var wg sync.WaitGroup
	wg.Add(5) //добавим группу
	for i := 0; i < 5; i++ {
		go func(j int) { //передаем в качестве параметра чтоб невызвать замыкание
			fmt.Println(j)
			wg.Done() // заверишим
		}(i)
	}
	wg.Wait() //ждем выполнения всех груп
}

func main7() {

	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины

	work := func(id int) {
		defer wg.Done() //Чтобы сигнализировать, что элемент группы завершил свое выполнение,
		fmt.Printf("Группа %d start", id)
		time.Sleep(2)
		fmt.Printf("Группа %d done", id)
	}
	// количество горутин, которые вызывают метод wg.Done() должно соответствовать количеству элементов группы wg
	go work(1)
	go work(2)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Print("DONE END")
}

func ParalelWrite(byt []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err := f1.Write(byt)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err := f1.Write(byt)
			res <- err
			f2.Close()
		}()
	}
	return res
}
