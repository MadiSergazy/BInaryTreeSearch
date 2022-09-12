package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
	len  int
}

// BST String (in o rder Traversal ) in codes        //СОРТИРОВКА
func (n Node) String() string {
	return strconv.Itoa(n.value) //int -> string
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}
func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *Node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))

	b.inOrderTraversalByNode(sb, root.right)
}

//

func main() {
	n := &Node{1, nil, nil}
	n.left = &Node{0, nil, nil}
	n.right = &Node{2, nil, nil}
	b := bst{
		root: n,
		len:  1,
	}
	fmt.Println(b)
}
