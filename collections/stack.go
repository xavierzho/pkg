package collections

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (v T) {
	t := s.Peek()
	// Remove the last element from the stack.
	*s = (*s)[:len(*s)-1]
	return t
}
func (s *Stack[T]) Peek() (v T) {
	if len(*s) == 0 {
		return
	}

	// Get the last element from the stack.
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}

func (s *Stack[T]) Values() []T {

	return *s
}
