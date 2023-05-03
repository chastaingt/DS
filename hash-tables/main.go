package main

import "fmt"

const ARRAY_SIZE = 7

type hashTable struct {
	array [ARRAY_SIZE]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func initHashTable() *hashTable {
	result := &hashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ARRAY_SIZE
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(key, "already exists")
	}
}

func (b *bucket) search(key string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}

func (h *hashTable) insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *hashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func main() {
	myHashTable := initHashTable()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		myHashTable.insert(v)
	}

	fmt.Println(myHashTable.search("KENNY"))
	myHashTable.insert("KENNY")
	myHashTable.delete("STAN")
	fmt.Println(myHashTable.search("STAN"))
}
