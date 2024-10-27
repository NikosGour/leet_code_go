package src

type Stack[T any] struct {
	stack []T
}

func (this *Stack[T]) push(v T) {
	this.stack = append(this.stack, v)
}

func (this *Stack[T]) pop() (T, error) {
	size := len(this.stack)
	if size == 0 {
		return *new(T), nil
	}
	res := this.stack[size-1]
	this.stack = this.stack[:size-1]
	return res, nil
}

func NewStack[T any]() Stack[T] {
	arr := make([]T, 0)
	return Stack[T]{stack: arr}
}

var close_to_open = map[rune]rune{')': '(', ']': '[', '}': '{'}

func isValid(s string) bool {

	stack := NewStack[rune]()

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.push(c)

		case ')', ']', '}':
			x, err := stack.pop()

			if err != nil {
				return false
			}

			if x != close_to_open[c] {
				return false
			}

		}
	}

	if len(stack.stack) != 0 {
		return false
	}
	return true

}
