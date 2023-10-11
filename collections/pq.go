package collections

import "container/heap"

type Item[T any] struct {
	v        T
	priority int
	idx      int
}
type PriorityQueue[T any] []*Item[T]

func NewPriorityQueue[T comparable](m map[T]int) PriorityQueue[T] {
	pq := make(PriorityQueue[T], len(m))
	i := 0
	for k, p := range m {
		pq[i] = &Item[T]{
			v:        k,
			priority: p,
			idx:      i,
		}
		i++
	}
	heap.Init(&pq)
	return pq
}
func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].idx = i
	(*pq)[j].idx = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.idx = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.idx = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.v = value
	item.priority = priority
	heap.Fix(pq, item.idx)
}
