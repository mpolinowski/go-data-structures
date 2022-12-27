package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Elements []int
}

// check stack size
func (s *Stack) Length() int {
	return len(s.Elements)
}

// stack is empty?
func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

// add an element to top of the stack
func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

// remove top element from stack
func (s *Stack) Pop() (int, error) {
	// check if stack is empty
	if s.IsEmpty() {
		return 0, errors.New("Stack is Empty")
	} else {
		var lastElement int
		lastElementIndex := len(s.Elements) - 1
		lastElement, s.Elements = s.Elements[lastElementIndex], s.Elements[:lastElementIndex]
		return lastElement, nil
	}
}

// check top element without removing
func (s *Stack) Peek() (int, error) {
	// check if stack is empty
	if s.IsEmpty() {
		return 0, errors.New("Stack is Empty")
	} else {
		return s.Elements[len(s.Elements)-1], nil
	}
}

func main() {
	fmt.Println("Stacks as Data Structures")

	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	elem1, _ := stack.Pop()
	fmt.Println(elem1)
	elem2, _ := stack.Pop()
	fmt.Println(elem2)

	peek3, _ := stack.Peek()
	fmt.Println(peek3)

	elem3, _ := stack.Pop()
	fmt.Println(elem3)

	fmt.Println(stack.IsEmpty())
}
