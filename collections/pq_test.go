package collections

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	items = nil
	pq := NewPriorityQueue[string](items)

	// Insert a new item and then modify its priority.
	item := &Item[string]{
		v:        "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.v, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[string])
		fmt.Printf("%.2d:%s \n", item.priority, item.v)
	}
}
