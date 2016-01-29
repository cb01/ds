package main


import (
	"testing"
)

func TestSet(t *testing.T) {

	h := NewHash()

	h.Set("hello", "kitty")

	k := h.hash("hello")
	if (h.data[k]).data[1] != "kitty" {
		t.Errorf("hash set error: value was never inserted in the hash")
	}

}

func TestGet(t *testing.T) {
	
	h := NewHash()

	h.Set("hello", "kitty")
	h.Set("hello", "world")

	if _, e := h.Get("kitty"); e == nil {
		t.Errorf("hash get errror: Hash.Get() should have returned an error when accessing non-existent key")
	}

	if v, _ := h.Get("hello"); v != "world" {
		t.Errorf("hash get error: Hash.Get() should have returned value 'world' for key 'hello', instead at that position was: %s", (*h.data[h.hash("hello")]).data[1])
	}

}

func TestDelete(t *testing.T) {

	h := NewHash()

	h.Set("hello", "kitty")
	h.Set("hi", "world")
	h.Delete("hello")

	if _, e := h.Get("hello"); e == nil {
		t.Errorf("hash delete error: Hash.Get() should have returned an error when accessing non-existent key, following deletion of that key")
	}

}

/*
func TestClustering(t *testing.T) {

	// all: shift to position
	// turn on: bit or 1, 0 -> 1, 1 -> 1
	// toggle: bit xor 1, 1 -> 0, 0 -> 1
	// turn off: bit and not(0), 1 -> 0, 0 -> 0
	// query: bit and 1, result & 0 implies value was 1

	numKeys := 20
	numBuckets := 49
	rsa := randStringArray(numKeys, 20)
	hkeys := make([]int, numKeys)

	for i, s := range rsa {
		k := hash(s, numBuckets)
		hkeys[i] = k
	}

	cscore := scoreHashClustering(hkeys, numBuckets)
	fmt.Printf("clustering of %d elements with %d buckets had score %f\n", numKeys, numBuckets, cscore)

}
*/


