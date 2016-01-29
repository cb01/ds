package main

import (
	"testing"
)

func TestStack(t *testing.T) {

	key := []int{1,4,2,6}	
	s := NewStack()

	if _, e := s.Pop(); e == nil {
		t.Errorf("Pop from empty stack should have returned non-nil error.")
	}

	for _, k := range key {
		e := s.Push(k)
		if e != nil {
			t.Errorf("%s",e)
		}
	}

	for i := 1; i <= len(key); i++ {

		ind := len(key) - i

		val, e := s.Pop()
		if e != nil {
			t.Errorf("%s",e)
		}
		if val != key[ind] {
			t.Errorf("error, value returned (%d) when popping from stack did not match what was expected (%d)", val, key[ind])
		}

	}

}
