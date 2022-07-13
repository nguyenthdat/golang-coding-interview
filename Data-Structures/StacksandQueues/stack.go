package main

import "fmt"

// stack
type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	temp := s.items[l]
	s.items = s.items[:l]
	return temp
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(33)
	myStack.Push(99)
	myStack.Pop()
	fmt.Println(myStack)
}
