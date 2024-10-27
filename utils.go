package main

type Stack[T any] struct {
	stack []T
}

func (this *Stack[T]) push(v T) {
	this.stack = append(this.stack, v)
}

func (this *Stack[T]) pop() T {
	size := len(this.stack)
	res := this.stack[size-1]
	this.stack = this.stack[:size-1]
	return res
}
