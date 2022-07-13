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

func (l *LinkedList) append(n *Node) {
	currentNode := l.head
	for i := 0; i < l.length; i++ {
		if currentNode.next == nil {
			currentNode.next = n
			return
		}
		currentNode = currentNode.next
	}
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

func (l *LinkedList) reverse() {
	if l.head.next == nil {
		return
	}
	firstNode := l.head
	secondNode := l.head.next
	for secondNode != nil {
		tempNode := secondNode.next
		secondNode.next = firstNode
		firstNode = secondNode
		secondNode = tempNode
	}
	l.head.next = nil
	l.head = firstNode
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
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}

	n4 := &Node{data: 4}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.append(n4)
	// myList.delete(32)
	myList.printLinkedListData()
	myList.reverse()
	fmt.Println()
	myList.printLinkedListData()
}
