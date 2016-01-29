package main 

import (
	"testing"
	//"fmt"
	"reflect"
)

// TODO: Needs more test coverage.

func TestPop(t *testing.T) {

	h := NewHeap(6)

	if _, e := h.Pop(); e == nil {
		t.Errorf("pop from empty heap should have returned non-nil error")
	}

	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)

	v, e := h.Pop()
	if e != nil {
		t.Errorf("%s", e)
	}
	if v != 5 {
		t.Errorf("error when popping from heap, expected a value of 5 and got %d", v)
	}

}

func TestComputeParentChildMappings(t *testing.T) {

	h := NewHeap(3)

	pcmapKey := map[int]map[int]int{
		0: map[int]int{
			1: 1,
			2: 2,
		},
		1: map[int]int{
			0: 0,
			1: 3,
			2: 4,
		},
		2: map[int]int{
			0: 0,
			1: 5,
			2: 6,
		},
		3: map[int]int{
			0: 1,
		},
		4: map[int]int{
			0: 1,
		},
		5: map[int]int{
			0: 2,
		},
		6: map[int]int{
			0: 2,
		},
	}

	if !reflect.DeepEqual(pcmapKey, h.pcmap) {
		t.Errorf("heap test: the computed parent/child index map did not match the testing key map")
	}

}

/*
func TestMaxHeapify(t *testing.T) {

	h := NewHeap(6)

	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)

	fmt.Println(h.data)

	h.data[0] = 4
	h.data[1] = 5
	h.data[2] = 3
	h.data[3] = 2
	h.data[4] = 1

	fmt.Println(h.data)

	h.maxHeapify(0)

	fmt.Println(h.data)


}
*/


