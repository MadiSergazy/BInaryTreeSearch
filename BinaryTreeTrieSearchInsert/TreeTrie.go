package main

import "fmt"

// Node

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node //будет указывать на начало структуры
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result

}

// insert
func (t *Trie) Insert(w string) {
	wordLenght := len(w)
	currentNode := t.root

	for i := 0; i < wordLenght; i++ {
		var charIndex byte
		if w[i] >= 'A' && w[i] <= 'Z' {
			charIndex = w[i] - 'A'
		} else {
			charIndex = w[i] - 'a' // -97 чтобы получить позицию в алфавите
		}

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex] //переходим на след
	}
	currentNode.isEnd = true
}

// search
func (t *Trie) search(w string) bool {

	wordLenght := len(w)
	currendNode := t.root

	for i := 0; i < (wordLenght); i++ {
		var charIndex byte
		if w[i] >= 'A' && w[i] <= 'Z' {
			charIndex = w[i] - 'A'
		} else {
			charIndex = w[i] - 'a'
		}

		if currendNode.children[charIndex] == nil {
			return false
		}
		currendNode = currendNode.children[charIndex]
	}

	if currendNode.isEnd == true {
		return true
	} else {
		return false
	}
}

func main() {

	toAdd := []string{
		"Madi",
		"Aryn",
		"Ayan",
		"Alkei",
		"Didar",
		"Bota",
		"Aldik",
		"Alesh",
	}

	testTrie := InitTrie()
	fmt.Println(testTrie)
	fmt.Println(testTrie.root)
	testTrie.Insert("Niger")
	fmt.Println(testTrie.search("Niger"))

	for _, v := range toAdd {
		testTrie.Insert(v)
	}
	fmt.Println(testTrie.search("Madih"))
}
