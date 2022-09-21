package main

import "fmt"

func main() {

	/*      //1 //Иди в зад
	fmt.Println("Я первый")
	word := "Message"
	//say(word)
	go say(word) //go создасть легковестный поток которыйм будет сам упарвлять
	fmt.Println("мы будем работаь пока say спит")
	fmt.Println("мы будем работаь пока say спит")
	fmt.Println("мы будем работаь пока say спиn")
	fmt.Println("мы будем работаь пока say спит") //почему это несработает потомушто когда основной поток завершится то завершится и горутины
	time.Sleep(time.Second * 2)
	//как передоватьданные из основного потока в горутины и из гоурутин в основной поток ответ КАНАЛЫ
	//ch := make(chan int)
	*/

	/*   //2
	ch := make(chan int)
	go sayHello(ch) //в другом горутине должно запуститься иначе будет панила

		//fmt.Println(<-ch) // <- читаем из канала
		//fmt.Println(<-ch) // <- читаем из канала
		//fmt.Println(<-ch) // <- читаем из канала
		//fmt.Println(<-ch) // <- читаем из канала
		//fmt.Println(<-ch) // <- читаем из канала

	for i := range ch {
		fmt.Println(i) //вместо этого итерируемся по каналу
	}

	//уже 6 раз деадлайн fmt.Println(<-ch) // <- читаем из канала
	*/

	data := make(chan int)
	exit := make(chan int)

	go func() { //анонимная функция
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0 //даем 0 в канал выхода
	}()
	select1(data, exit)
}

/*         //1
func say(word string) {
	time.Sleep(2 * time.Second)
	fmt.Println(word)
	fmt.Println("Нигер нигери ")
}
*/

/*           //2
func say(s string) {
	fmt.Println(s)
}

func sayHello(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		say("HELLO")
		ch <- i //запис в канал
	}
	close(ch) //закрываем канал чтобы небыло ошибки чтобы неждали запис
}
*/

func select1(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("Exit")
			return
		}
	}
}
