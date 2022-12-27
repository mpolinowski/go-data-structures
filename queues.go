package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
}

// show length of queue
func (q *Queue) Length() int {
	return len(q.Elements)
}

// return true when queue length is zero
func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

// add element to end of queue
func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

// check first element in que without removing it
func (q *Queue) Peek() (int, error) {
	// test queue not empty
	if q.IsEmpty() {
		return 0, errors.New("Empty Queue")
	}
	return q.Elements[0], nil
}

// remove first element from queue
func (q *Queue) Dequeue() (int, error) {
	// test queue not empty
	if q.IsEmpty() {
		return 0, errors.New("Empty Queue")
	}
	// drop first element from slice
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

func main() {
	fmt.Println("Queues as Data Structures")

	// instantiate queue and add element
	var elem int
	queue := Queue{}
	fmt.Println(queue)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue)

	// drop first item
	elem, _ = queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

	// check first item without removing it
	elem, _ = queue.Peek()
	fmt.Println(elem)
	fmt.Println(queue)
}
