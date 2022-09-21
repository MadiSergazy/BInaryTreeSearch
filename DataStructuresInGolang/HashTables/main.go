package main

import "fmt"

const ArraySize = 7 // размер хэш таблицы

// HashTable struct
type HachTable struct { //will hold on array
	array [ArraySize]*bucket
}

// bucked struct is a linked list in each slot of the Hach Table array
type bucket struct {
	head *bucketNode
}

// BuckedNode struct
// is a linkedlist node that holds the key
type bucketNode struct {
	key  string      // names
	next *bucketNode //addres on next node
}

// insert HashTable
// will take akey and add it to hachTable
func (h *HachTable) Insert(key string) {
	index := hach(key)
	h.array[index].insert(key)
}

// search HashTable
func (h *HachTable) Search(key string) bool {
	index := hach(key)
	return h.array[index].search(key)
}

// delete HashTable
func (h *HachTable) Delete(key string) {
	index := hach(key)
	h.array[index].delete(key)
}

// insert Bucked will take akeycreate a nodewith the key and insert thenode in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Already Exists")
	}

}

// search Buckedwill take a key and return true if bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil { //пробегаемся по ведру если найдем вернем тру
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false

}

// delete Bucked will take a key and delete node from bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
	fmt.Println("Zdes net takogo key")
}

// Hash
func hach(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
	/*
		sum := byte(0)
		counter := byte(0)
		for _, k := range key {
			sum += byte(k)
			counter++
		}
		return int(sum % counter)*/
}

// Init - it create bucketin each slot of hach table
func Init() *HachTable {
	result := &HachTable{}
	for i := range result.array {
		result.array[i] = &bucket{} //внутри каждого элемента массива создадим ведро для связанных  списков
	}
	return result
}
func main() {
	//testHachTable := &HachTable{}
	testHachTable := Init()
	fmt.Println(testHachTable)
	fmt.Println(hach("Randy"))

	testBucked := &bucket{}
	testBucked.insert("Rabdy")
	testBucked.insert("Eric")

	fmt.Println(testBucked.search("Eric"))
	fmt.Println(testBucked.search("Randy"))

	testBucked.delete("Randy")
	fmt.Println(testBucked.search("Randy"))
}
