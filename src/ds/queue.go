package main 

import (
	"fmt"
)

// Queue is an implementation of a classic FIFO queue data
// structure.
type Queue struct {
	head *qNode
	tail *qNode
}

type qNode struct {
	next *qNode
	data int
}

// NewQueue creates a new empty Queue.
func NewQueue() Queue {
	return Queue{}
}

// Enqueue enqueues an integer value into a Queue.
//
// Enqueue will return a non-nil error if a failure occurs 
// when attempting to enqueue the value.
func (q *Queue) Enqueue(val int) error {

	// Create a new node to store the data that is being inserted
	// into the queue
	n := qNode{nil, val}

	// If the queue is empty, point to this node as both the head 
	// and tail and return.
	if q.head == nil {
		q.head = &n
		q.tail = &n
		return nil
	}

	// The queue is non-empty so attach this node to the tail node
	// and update the tail of the queue.
	(*q.tail).next = &n
	q.tail = &n

	return nil
}

// Dequeue returns the value at the front of the queue.
//
// If Dequeue() is called on an empty queue, a non-nil error
// will be returned.
func (q *Queue) Dequeue() (int, error) {
	
	if q.head == nil {
		return -1, fmt.Errorf("Can't return a value from an empty queue.")
	}

	// Obtain the value of the first element in the queue (which
	// will be returned)
	val := (*q.head).data

	// If the head node is also the last node in the queue, set both
	// the head and tail pointers to nil
	if (*q.head).next == nil {
		q.head = nil
		q.tail = nil
		return val, nil
	}

	// Otherwise, update only the head pointer
	q.head = (*q.head).next

	return val, nil
}

// ChanQueue is a queue implemented using a channel
type ChanQueue struct {
	size int
	stream chan int
}

// Enqueue enqueues a value into ChanQueue
func (cq *ChanQueue) Enqueue(val int) error {

	// A value is enqueued by writing to the ChanQueue.stream channel
	// and noting that the size has increased by 1.
	cq.size += 1
	cq.stream <- val
	return nil

}

// Dequeue reads a value from the ChanQueue queue
//
// A non-nil error will be returned if the queue is already empty.
func (cq *ChanQueue) Dequeue() (int, error) {

	if cq.size > 0 {
		cq.size -= 1
		val := <- cq.stream
		return val, nil
	} else {
		return 0, fmt.Errorf("queue: cannot return a value from an empty queue")
	}

}


