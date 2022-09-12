package BinaryTreeSearch

import "fmt"

//Node

var steps int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// insert method
func (n *Node) insert(k int) { //(n *Node)  - указываем что это метод структуры
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

// search method
func (n *Node) search(k int) bool {
	/*if n.Key == k {
		return true
	} else if n.Key < k {
		if n.Right.Key == k {
			return true
		}
		if n.Right.Key < k {
			n.Right.searhc(k)
		} else {
			n.Left.searhc(k)
		}
	}*/
	steps++
	if n == nil { //если такого элемента нету
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.search(k)
	} else if n.Key > k {
		return n.Left.search(k)
	} else {
		return true
	}

}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)

	tree.insert(155)
	fmt.Println(tree)

	tree.insert(55)
	fmt.Println(tree)

	tree.insert(25)
	tree.insert(75)
	fmt.Println(tree)

	fmt.Println(tree.search(75))
	fmt.Println(steps)
}
