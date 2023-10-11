package collections

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap[int](func(a, b int) bool { return a < b })
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	for h.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(h))
	}
}
