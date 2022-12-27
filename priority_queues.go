package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	High []int
	Low  []int
}

// show length of complete queue
func (q *PriorityQueue) Length() int {
	return len(q.High) + len(q.Low)
}

// return true when queue length is zero
func (q *PriorityQueue) IsEmpty() bool {
	return q.Length() == 0
}

// add element to end of queue
// differentiate between high/low priority
func (q *PriorityQueue) Enqueue(elem int, priority bool) {
	if priority {
		q.High = append(q.High, elem)
	} else {
		q.Low = append(q.Low, elem)
	}
}

// check first element in que without removing it
// but prefer priority queue
func (q *PriorityQueue) Peek() (int, error) {
	// test queue not empty
	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}
	// return 0 if both queues are empty
	return 0, errors.New("Empty Queues")
}

// remove first element from queue
// but prefer priority queue
func (q *PriorityQueue) Dequeue() (int, error) {
	// test if priority queue not empty
	// if true remove first element
	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}
	// test if regular queue not empty
	// if true remove first element
	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}
	// return 0 if both queues are empty
	return 0, errors.New("Empty Queues")
}

func main() {
	fmt.Println("Priority Queues as Data Structures")

	// instantiate queue and add element
	var elem int
	queue := PriorityQueue{}
	fmt.Println(queue)
	queue.Enqueue(1, true)
	queue.Enqueue(2, true)
	queue.Enqueue(3, false)
	queue.Enqueue(4, false)
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
