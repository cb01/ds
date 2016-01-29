package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	
	key := []int{1,3,5,7,9}

	q := NewQueue()

	if _, e := q.Dequeue(); e == nil {
		t.Errorf("Queue.Dequeue() should have given error on dequeue from empty queue.")
	}

	for _, v := range key {
		q.Enqueue(v)
	}

	for _, kv := range key {
		v, e := q.Dequeue()
		if e != nil {
			return
		}
		if v != kv {
			t.Errorf("test error: value returned (%d) when dequeueing does not match expected value (%d) from key used to create queue", v, kv)
		}
	}

}


