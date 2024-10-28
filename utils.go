package main

type Stack[T any] struct {
	array []T
}

func (this *Stack[T]) push(v T) {
	this.array = append(this.array, v)
}

func (this *Stack[T]) pop() T {
	size := len(this.array)
	res := this.array[size-1]
	this.array = this.array[:size-1]
	return res
}

func NewStack[T any]() Stack[T] {
	arr := make([]T, 0)
	return Stack[T]{array: arr}
}

type Set[T comparable] struct {
	_map map[T]bool
}

func (this *Set[T]) append(val T) {
	if !this.contains(val) {
		this._map[val] = true
	}
}
func (this *Set[T]) contains(val T) bool {
	_, exist := this._map[val]
	return exist
}

func NewSet[T comparable]() Set[T] {
	_map := make(map[T]bool)
	return Set[T]{_map: _map}
}
