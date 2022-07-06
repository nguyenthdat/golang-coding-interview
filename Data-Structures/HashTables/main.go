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
// func (h *HashTable) Search(key string) bool {
// 	index := Hash(key)
// }

// // Delete
// func (h *HashTable) Delete(key string) {
// 	index := Hash(key)
// }

// insert
func (b *Bucket) insert(key string) {
	newNode := &BucketNode{key: key}
	newNode.next = b.head
	b.head = newNode
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
	test := Init()
	fmt.Println(test)
	fmt.Println(Hash("TEST"))

	testB := &Bucket{}
	test.Insert("TEST2")
	fmt.Println(testB)
}
