package stack

import (
	"fmt"
	"strings"
)

// EmptyStackTopIndex 当栈为空的时候栈顶的指针指向的下标位置
const EmptyStackTopIndex = -1

// ArrayStack 数组实现的栈
type ArrayStack[T any] struct {

	// 栈顶指针，始终指向栈顶元素所在的下标
	topIndex int

	// 栈中的元素，栈底是0，会不断往上弹
	slice []T
}

var _ Stack[any] = &ArrayStack[any]{}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{
		slice:    make([]T, 0),
		topIndex: EmptyStackTopIndex,
	}
}

func (x *ArrayStack[T]) Push(values ...T) {
	for _, value := range values {
		nextIndex := x.topIndex + 1
		if nextIndex < len(x.slice) {
			x.slice[nextIndex] = value
		} else {
			x.slice = append(x.slice, value)
		}
		x.topIndex = nextIndex
	}
}

func (x *ArrayStack[T]) Pop() T {
	e, _ := x.PopE()
	return e
}

func (x *ArrayStack[T]) PopE() (T, error) {
	if x.topIndex == EmptyStackTopIndex {
		var zero T
		return zero, ErrStackEmpty
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
	if x.topIndex == EmptyStackTopIndex {
		var zero T
		return zero, ErrStackEmpty
	}
	return x.slice[x.topIndex], nil
}

func (x *ArrayStack[T]) IsEmpty() bool {
	return x.topIndex == EmptyStackTopIndex
}

func (x *ArrayStack[T]) IsNotEmpty() bool {
	return x.topIndex != EmptyStackTopIndex
}

func (x *ArrayStack[T]) Size() int {
	return x.topIndex + 1
}

func (x *ArrayStack[T]) Clear() {
	x.topIndex = EmptyStackTopIndex
	x.slice = make([]T, 0)
}

func (x *ArrayStack[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	currentIndex := x.topIndex
	for currentIndex != EmptyStackTopIndex {
		sb.WriteString(fmt.Sprintf("%#v", x.slice[currentIndex]))
		currentIndex--
	}
	sb.WriteString("]")
	return sb.String()
}
