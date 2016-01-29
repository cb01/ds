package main 

import (
	"fmt"
)

// A heap is a balanced binary tree data structure where the value
// of each node is larger than that of either of its children. In
// this case, we're describing a max heap.
type Heap struct {

	// An array of integers with data[0] corresponding to the root of the
	// tree and data[len(data)-1] corresponding to the rightmost leaf. The
	// two elements at level 1 are at positions data[1] and data[2], the
	// four at level 2 at data[3], data[4], .. data[6], and so on.
	data []int

	// The integer index of the last member of the array
	last int

	// The current size of the allocated array
	allocatedSize int

	// pcmap stores a mapping from any index to the indices of the parent
	// and child indices of that node, if they exist.
	pcmap map[int]map[int]int

}

// NewHeap creates and returns a new empty heap data structure
func NewHeap(allocatedSize int) Heap {
	a := make([]int, allocatedSize)
	pcmap := map[int]map[int]int{0:map[int]int{0:0}}
	h := Heap{a, 0, allocatedSize, pcmap}
	h.computeParentChildMappings()
	return h
}

// computeParentChildMappings pre-computes the indices of child and parent
// nodes for each node in the heap, ideally allowing O(1) lookup of heap
// elements either by id or by association.
func (h *Heap) computeParentChildMappings() {

	// Pre-compute the indices of parents and children up to a number
	// that is relevant given the current size of the heap. pcmap[i][0]
	// corresponds to the parent of node i, with pcmap[i][1] corresponding
	// to the left child and pcmap[i][2] corresponding to the right child.
	lastLevelStart := 0
	lastLevelLength := 1
	total := 0
	h.pcmap[0] = map[int]int{}

	for {

		// For each node in the previous level, link it to the correct
		// one in the current level
		llend := lastLevelStart + lastLevelLength
		for i := 0; i < lastLevelLength; i++ {
			pInd := lastLevelStart + i
			cLeft := llend + i*2
			cRight := llend + 2*i + 1
			h.pcmap[pInd][1] = cLeft // mark node at pInd and having left child at cLeft
			h.pcmap[pInd][2] = cRight // mark node at pInd as having right child at cRight
			
			// Making reference to child nodes for first time to allocate them in the map
			h.pcmap[cLeft] = map[int]int{}
			h.pcmap[cLeft][0] = pInd // mark node at cLeft as having parent at pInd
			h.pcmap[cRight] = map[int]int{}
			h.pcmap[cRight][0] = pInd // mark node at cRight as having parent at pInd
			total += 1
			// If we've computed enough parent/child mappings to have one for
			// each node in the currently allocated heap, we are done.
			if total >= h.allocatedSize {
				break
			}
		}

		if total >= h.allocatedSize {
			break
		}

		lastLevelStart += lastLevelLength
		lastLevelLength = lastLevelLength*2

	}

}

// Pop returns the maximum value of a heap, deleting it and 
// restoring the max heap property of the heap.
//
// Determining the maximum value is O(1) but removing it and restoring
// the max heap property may take log(n) operations, therefore this
// is O(log(n)).
func (h *Heap) Pop() (int, error) {
	
	if h.last == 0 {
		return -1, fmt.Errorf("heap: cannot pop from empty heap")
	}

	// Obtain the largest value from the heap which will be the first 
	// element in the underlying array.
	val := h.data[0]

	// Move the rightmost leaf of the heap to the root and decrement the
	// index of the last element in the array.
	h.data[0] = h.data[h.last]
	h.last -= 1

	// Make a call to the recursive MaxHeapify function, starting at the
	// root node.
	h.maxHeapify(0)

	return val, nil
}

// leftChild returns the left child of a node in a heap at index ind.
//
// If the node does not have a left child, a non-nil error is returned.
func (h *Heap) leftChild(ind int) (int, error) {

	if _, ok := h.pcmap[ind]; !ok {
		fmt.Errorf("heap: tried to access left child for a node not present in the precomputed parent/child index map")
	}

	cInd, ok := h.pcmap[ind][1]
	if !ok {
		return -1, fmt.Errorf("heap: tried to access right child for indexed node for which the index of a left child node has not been computed")
	}

	if cInd > h.last {
		return -1, fmt.Errorf("heap: node does not have left child")
	}

	return cInd, nil
}

// rightChild returns the right child of a node in a heap at index ind.
//
// If the node does not have a right child, a non-nil error is returned.
func (h *Heap) rightChild(ind int) (int, error) {

	if _, ok := h.pcmap[ind]; !ok {
		return -1, fmt.Errorf("heap: tried to access right child for a node not present in the precomputed parent/child index map")
	}

	cInd, ok := h.pcmap[ind][2]
	if !ok {
		return -1, fmt.Errorf("heap: tried to access right child for indexed node for which the index of a right child node has not been computed")
	}

	if cInd > h.last {
		return -1, fmt.Errorf("heap: node does not have right child")
	}

	return cInd, nil
}

// parent computes the index of the parent of a node in a max heap.
//
// parent will return an error if the node at the specified index does not
// have a parent
func (h *Heap) parent(ind int) (int, error) {

	if _, ok := h.pcmap[ind]; !ok {
		return -1, fmt.Errorf("heap: tried to access parent for a node not present in the precomputed parent/child index map")
	}

	if ind == 0 {
		return -1, fmt.Errorf("heap: node does not have a parent")
	}

	pInd, _ := h.pcmap[ind][0]
	//if !ok {
	//	return -1, fmt.Errorf("heap: node does not have parent despite having an index greater than 1, problem.")
	//}

	return pInd, nil
}

// MaxHeapify recursively restores the MaxHeap property of a heap, beginning
// at a specified index.
func (h *Heap) maxHeapify(ind int) {

	// The value of the data stored at the current node being evaluated
	currentNodeData := h.data[ind]

	// If the value at the provided index is less than its children,
	// bubble down.

	lcInd, el := h.leftChild(ind);
	rcInd, er := h.rightChild(ind);
	pInd, ep := h.parent(ind);

	if (el == nil && h.data[lcInd] > currentNodeData) && ((er == nil && (h.data[lcInd] >= h.data[rcInd])) || (er != nil)) {
		h.data[ind] = h.data[lcInd]
		h.data[lcInd] = currentNodeData
		h.maxHeapify(lcInd)
	} else if (er == nil && h.data[rcInd] > currentNodeData) {
		h.data[ind] = h.data[rcInd]
		h.data[rcInd] = currentNodeData
		h.maxHeapify(rcInd)
	} else if (ep == nil && h.data[pInd] < currentNodeData) {
		// Otherwise, if the value at the provided index is greater than its parent,
		// bubble up.
		h.data[ind] = h.data[pInd]
		h.data[pInd] = currentNodeData
		h.maxHeapify(pInd)
	}

	// Otherwise, the value at the provided index is max-heap correct, return
	return

}

// Insert inserts a new value into the heap, beginning at a leaf node
// and swapping upwards toward the root node until the inserted node
// has a value less than its parent.
//
// O(log(n)): Always balanced binary tree, therefore the height is 
// log(n). Wost case is that each new node is the largest observed
// thus far and must be compared to every node from that leaf to the
// root, which is equal to the height of the tree, or log(n).
func (h *Heap) Insert(val int) error {
	
	// If the heap does not have enough space to insert a new value, 
	// increase its size by a factor of k.
	if h.last + 1 >= h.allocatedSize {
		newData := make([]int, 0, h.allocatedSize*2)
		copy(newData, h.data)
		h.data = newData
	}

	// Insert the new value as the last element in the heap array.
	h.data[h.last] = val

	// Call max-heapify to restore the max heap property of the heap.
	h.maxHeapify(h.last)

	h.last += 1
	return nil
}


