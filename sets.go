package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Elements map[string]struct{}
}

// create new set
func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// add elements to set
// sets don't have an order - we don't need to append
// every element in a set is unique - already existing
// elements will not be duplicated (no checks necessary)
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

// remove element from set if exists
func (s *Set) Delete(elem string) error {
	// check if element is present
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("Element not present!")
	}
	// if present -delete
	delete(s.Elements, elem)
	return nil
}

// check if set contains element
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

// list all elements from set
func (s *Set) List() {
	for key, _ := range s.Elements {
		fmt.Println(key)
	}
}

func main() {
	fmt.Println("Sets as Data Structures")

	// instantiate / populate set
	mySet := NewSet()
	mySet.Add("Eddard Stark")
	mySet.Add("Jaime Lannister")
	mySet.Add("Daenerys Targaryen")
	mySet.Add("Arya Stark")
	mySet.Add("Sandor Clegane")
	mySet.Add("Tyrion Lannister")

	// delete element
	mySet.Delete("Eddard Stark")

	// find element
	fmt.Println(mySet.Contains("Jon Snow"))
	fmt.Println(mySet.Contains("Arya Stark"))

	// list all elements
	mySet.List()
}
