/*
type Node struct {
	Key   float64
	Left  *Node
	Right *Node
}

// insert method
func (n *Node) insert(k float64) { //(n *Node)  - указываем что это метод структуры
	if n.Key < k { //если вставляемое знач больше то  впарва
		if n.Right == nil { //если мето пусто
			n.Right = &Node{Key: k}
		} else {
			n.Right.insert(k) //если там есть значение то вызываем опять вставку
		}

	} else {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}
func main() {
	//fmt.Println(printHex("1434848"))
	tree := &Node{Key: 1}
	arguments := os.Args[1:]
	var mas []float64

	for _, v := range arguments {

		san, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		} else {

			mas = append(mas, san)
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(mas)))
	fmt.Println(mas)

	ans := make(map[int]string)
	var ans2 []string

	for i := len(mas); i >= 2; i-- {
		ans[i-1] += "1"
		ans[i-2] += "0"

		ans2 = append(ans2, "0")
		ans2 = append(ans2, "1")
		//ans2[i-1] += "0"
		//ans2[i-2] += "1"

		apElem := mas[i-1] + mas[i-2]
		mas = mas[:i-2]
		mas = append(mas, apElem)
		sort.Sort(sort.Reverse(sort.Float64Slice(mas)))
		fmt.Println(mas)
		tree.insert(apElem)

	}

	fmt.Println(ans)
	fmt.Println(ans2)
	fmt.Println(tree)
}
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Elem struct {
	name string
	sum  float64
	sub  []Elem //container
	val  string
}

func main() {
	var mas []Elem
	args := os.Args[1:]
	var pyramid []float64
	var sum float64
	for _, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		}
		pyramid = append(pyramid, num)
		sum += num
	}
	if sum > 1 {
		fmt.Println("SUM OF MASSIV MORE THAN 1")
		return
	}

	for k, i := range pyramid { //veroyatnost
		var newEl Elem
		newEl.sum = i
		newEl.name = strconv.Itoa(k + 1)
		mas = append(mas, newEl)
	}
	mas = Sort(mas)
	for len(mas) > 1 { //kodirovanie
		var newEl Elem
		newEl.sum = mas[len(mas)-2].sum + mas[len(mas)-1].sum
		newEl.sub = append(newEl.sub, mas[len(mas)-2])
		newEl.sub = append(newEl.sub, mas[len(mas)-1])
		newEl.name = mas[len(mas)-2].name + mas[len(mas)-1].name
		mas = mas[:len(mas)-2]
		mas = append(mas, newEl)
		mas = Sort(mas)
		//	fmt.Println(mas)
	}
	print(mas[0])
}

func print(mas Elem) {
	if len(mas.sub) != 0 {
		mas.sub[0].val = mas.val + "1"
		mas.sub[1].val = mas.val + "0"
		print(mas.sub[0])
		print(mas.sub[1])
	} else {
		fmt.Println(mas.name, mas.sum, mas.val)
	}
}

func Sort(mass []Elem) []Elem {
	var mas []Elem
	mas = mass
	for i := len(mas) - 1; i >= 1; i-- { //sort
		for j := i - 1; j >= 0; j-- {
			if mas[i].sum > mas[j].sum {
				mas[i].val, mas[j].val = mas[j].val, mas[i].val
				mas[i].name, mas[j].name = mas[j].name, mas[i].name
				mas[i].sum, mas[j].sum = mas[j].sum, mas[i].sum
				mas[i].sub, mas[j].sub = mas[j].sub, mas[i].sub
			}
		}
	}
	return mas
}
