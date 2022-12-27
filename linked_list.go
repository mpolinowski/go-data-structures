package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Size int
}

// nodes are linked elements in list
type Node struct {
	Value string
	Next  *Node
}

// add new node to head of linked list
func (l *LinkedList) Insert(elem string) {
	node := Node{
		Value: elem,
		Next:  l.Head,
	}
	l.Head = &node
	l.Size++
}

// remove first element
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

// iterate through list and print
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

// find element in list
func (l *LinkedList) Search(elem string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == elem {
			return current
		}
		current = current.Next
	}
	return nil
}

// delete element
func (l *LinkedList) Delete(elem string) {
	previous := l.Head
	current := l.Head
	// check if element exists
	for current != nil {
		// link previous to next to remove current
		if current.Value == elem {
			previous.Next = current.Next
			l.Size--
		}
		previous = current
		current = current.Next
	}
}

func main() {
	fmt.Println("Linked Lists as Data Structures")
	var ll LinkedList

	ll.Insert("Camina Drummer")
	ll.Insert("Joe Miller")
	ll.Insert("Amos Burton")
	ll.Insert("Chrisjen Avasarala")

	ll.List()
	fmt.Println("-----------------------------")

	ll.DeleteFirst()
	ll.List()
	fmt.Println("-----------------------------")

	if element := ll.Search("Joe Miller"); element != nil {
		fmt.Printf("%+v\n", element)
	}
	fmt.Println("-----------------------------")

	ll.Delete("Joe Miller")
	ll.List()
	fmt.Println(ll.Size)
}
