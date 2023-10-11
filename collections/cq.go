package collections

import (
	"container/list"
	"errors"
)

type CircleQueue[T any] struct {
	entry []T
	max   int
	head  int
	tail  int
}

var (
	ErrFull  = errors.New("queue is full")
	ErrEmpty = errors.New("queue is empty")
)

func NewCircleQueue[T any](m int) *CircleQueue[T] {
	list.New()
	return &CircleQueue[T]{
		entry: make([]T, m, m),
		max:   m,
		head:  0,
		tail:  0,
	}
}

func (cq *CircleQueue[T]) Push(v T) error {
	if cq.IsFull() {
		return ErrFull
	}
	cq.entry[cq.tail] = v
	cq.tail = (cq.tail + 1) % cq.max
	return nil
}
func (cq *CircleQueue[T]) Pop() (v T, err error) {
	if cq.IsEmpty() {
		return v, ErrEmpty
	}
	v = cq.entry[cq.head]
	cq.head = (cq.head + 1) % cq.max
	return
}
func (cq *CircleQueue[T]) List() ([]T, error) {
	var size = cq.Size()
	var result = make([]T, 0, cq.max)

	if size == 0 {
		return result, ErrEmpty
	}
	idx := cq.head
	for i := 0; i < size; i++ {
		result = append(result, cq.entry[idx])
		idx = (idx + 1) % cq.max
	}
	return result, nil
}
func (cq *CircleQueue[T]) IsFull() bool {
	return (cq.tail+1)%cq.max == cq.head
}

func (cq *CircleQueue[T]) IsEmpty() bool {
	return cq.tail == cq.head
}

func (cq *CircleQueue[T]) Size() int {
	return (cq.tail + cq.max - cq.head) % cq.max
}
