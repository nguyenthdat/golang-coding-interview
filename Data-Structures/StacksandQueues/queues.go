package main

import "fmt"

// Queue
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue
func (q *Queue) Dequeue() int {
	temp := q.items[0]
	q.items = q.items[1:]
	return temp
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)
}
