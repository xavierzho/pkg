package collections

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// An Heap is a min-heap of ints.
type Heap[T any] struct {
	entry []T
	less  func(T, T) bool
}

func NewHeap[T any](less func(T, T) bool) *Heap[T] {
	var h Heap[T]
	h.entry = make([]T, 0)
	h.less = less
	return &h
}
func (h *Heap[T]) Len() int           { return len(h.entry) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.entry[i], h.entry[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.entry[i], h.entry[j] = h.entry[j], h.entry[i] }

func (h *Heap[T]) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.entry = append(h.entry, x.(T))
}

func (h *Heap[T]) Pop() interface{} {
	old := h.entry
	n := len(old)
	x := old[n-1]
	h.entry = old[0 : n-1]
	return x
}

func (h *Heap[T]) Peek() T {
	return h.entry[0]
}
