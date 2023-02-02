// Итоговое задание для проверки #2:

// Иногда проще оперировать словами, а не предложениями.
// Напишите конвейер, что принимает строки и разделяет их на слова
// (можете использовать функцию Fields из пакета strings), а также
// отправляет все слова, одно за другим, на следующую стадию конвейера.

package main

import (
	"fmt"
	"strings"
)

func main() {
	//mas := []string{"hello world", "a bad apple", "goodbye all"}
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1) //,// mas)
	go splitWords(c1, c2)
	go printGopher(c2)
}

func sourceGopher(downstream chan string) { //, //mas []string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {

		downstream <- v
	}
	close(downstream)
}

func splitWords(upstream, downstream chan string) {
	for v := range upstream {
		for _, word := range strings.Fields(v) {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
