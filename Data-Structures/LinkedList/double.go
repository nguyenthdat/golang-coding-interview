package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	}
	currentNode := l.head
	l.head = n
	l.head.next = currentNode
	l.length++
}

func (l *LinkedList) append(n *Node) {
	if l.head == nil || l.tail == nil {
		l.head = n
		l.tail = n
	}
	l.tail.next = n
	l.tail = n
	l.length++
}

func (l LinkedList) printLinkedListData() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
}

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

func main() {
	myList := LinkedList{}
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}

	n4 := &Node{data: 4}
	n5 := &Node{data: 5}

	myList.append(n4)

	myList.append(n1)
	myList.append(n2)
	myList.append(n3)
	myList.prepend(n5)
	// myList.delete(32)
	myList.printLinkedListData()
}
