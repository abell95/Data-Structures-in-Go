package main

import (
	"fmt"
)

type Stack struct {
	Head *StackItem	
	size int
}

type StackItem struct {
	value int
	previous *StackItem
}

func (s *Stack) push(val int) {
	s.Head = &StackItem{val, s.Head}
	s.size++
}

func (s *Stack) pop() int {
	if (s.Head != nil) {
		popped := s.Head.value
		s.Head = s.Head.previous
		s.size--
		return popped
	} else {
		fmt.Println("Cant pop empty stack ya DINGUS and this zero means nothing")
		return 0
	}
}

func (s Stack) isEmpty() bool {
	return s.Head == nil
}

func (s Stack) peek() int {
	if s.isEmpty() != true {
		return s.Head.value
	} else {
		fmt.Println("Cant peek empty stack")
		return 0
	}
}

func (s Stack) getSize() int {
	return s.size
}

func main() {
	stack := Stack{nil, 0}
	stack.push(5)
	fmt.Println(stack.pop())
	fmt.Println(stack.isEmpty())
	stack.push(6)
	stack.push(9)
	fmt.Println("Size is",stack.getSize())
}
