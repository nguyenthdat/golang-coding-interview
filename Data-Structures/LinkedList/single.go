package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// prepend
func (l *LinkedList) prepend(n *Node) {
	currentNode := l.head
	l.head = n
	l.head.next = currentNode
	l.length++
}

// search
// func (l *LinkedList) search(s int) {
// }

// delete
func (l *LinkedList) delete(data int) {
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	prevNode := l.head
	for prevNode.next != nil {
		if prevNode.next.data == data {
			// delete
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

// print all nodes
func (l LinkedList) printLinkedListData() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	myList := LinkedList{}
	n1 := &Node{data: 22}
	n2 := &Node{data: 32}
	n3 := &Node{data: 12}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.delete(32)
	myList.printLinkedListData()
}
