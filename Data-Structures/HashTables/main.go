package main

import "fmt"

const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket structure
type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert
func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.array[index].search(key)
}

// // Delete
func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].delete(key)
}

// insert
func (b *Bucket) insert(key string) {
	if !b.search(key) {
		newNode := &BucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("%v already exist\n", key)
	}
}

// search
func (b *Bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *Bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == key {
			// delete
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

// Hash
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"TEST1",
		"TEST2",
		"TEST3",
		"TEST4",
		"TEST5",
		"TEST6",
		"TEST7",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("TEST6")
	fmt.Println(hashTable.Search("TEST6"))
}
