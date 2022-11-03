package stack

import "fmt"

// ArrayStack 数组实现的栈
type ArrayStack[T any] struct {

	// 栈顶指针
	topIndex int

	// 栈中的元素，栈底是0，会不断往上弹
	slice []T
}

var _ Stack[any] = &ArrayStack[any]{}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{
		slice:    make([]T, 0),
		topIndex: 0,
	}
}

func (x *ArrayStack[T]) Push(values ...T) error {
	for _, value := range values {
		if x.topIndex < len(x.slice) {
			x.slice[x.topIndex] = value
		} else {
			x.slice = append(x.slice, value)
		}
		x.topIndex++
	}
	return nil
}

func (x *ArrayStack[T]) Pop() T {
	e, _ := x.PopE()
	return e
}

func (x *ArrayStack[T]) PopE() (T, error) {
	if x.topIndex == 0 {
		return nil, ErrStackEmpty
	}
	value := x.slice[x.topIndex]
	x.topIndex--
	return value, nil
}

func (x *ArrayStack[T]) Peek() T {
	e, _ := x.PeekE()
	return e
}

func (x *ArrayStack[T]) PeekE() (T, error) {
	if x.topIndex == 0 {
		return nil, ErrStackEmpty
	}
	return x.slice[x.topIndex], nil
}

func (x *ArrayStack[T]) IsEmpty() bool {
	return x.topIndex == 0
}

func (x *ArrayStack[T]) IsNotEmpty() bool {
	return x.topIndex != 0
}

func (x *ArrayStack[T]) Len() int {
	return x.topIndex
}

func (x *ArrayStack[T]) Clear() error {
	x.topIndex = 0
	x.slice = make([]T, 0)
	return nil
}

func (x *ArrayStack[T]) String() string {
	return fmt.Sprintf("%#v", x)
}
