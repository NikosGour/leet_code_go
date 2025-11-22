package internal

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) push(v T) {
	s.arr = append(s.arr, v)
}

func (s *Stack[T]) pop() T {
	size := len(s.arr)
	res := s.arr[size-1]
	s.arr = s.arr[:size-1]
	return res
}

func (s *Stack[T]) size() int {
	return len(s.arr)
}

func NewStack[T any]() Stack[T] {
	arr := make([]T, 0)
	return Stack[T]{arr: arr}
}
