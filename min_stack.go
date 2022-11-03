package stack

type MinStack[T any] struct {
	comparator Comparator[T]
	stack      Stack[*MinNode[T]]
}

var _ Stack[any] = &MinStack[any]{}

func NewMinStack[T any](comparator Comparator[T]) *MinStack[T] {
	return &MinStack[T]{
		comparator: comparator,
		stack:      NewArrayStack[*MinNode[T]](),
	}
}

func (x *MinStack[T]) Push(values ...T) {
	for _, value := range values {
		node := &MinNode[T]{
			value: value,
		}
		topNode, err := x.stack.PeekE()
		if err != nil {
			// 栈为空，认为最小值是自己
			node.min = value
		} else if x.comparator(topNode.min, value) < 0 {
			// 当前栈顶的元素比较小
			node.min = topNode.min
		} else {
			node.min = value
		}
		x.stack.Push(node)
	}
}

func (x *MinStack[T]) Pop() T {
	e, _ := x.PopE()
	return e
}

func (x *MinStack[T]) PopE() (T, error) {
	e, err := x.stack.PopE()
	if err != nil {
		var zero T
		return zero, err
	}
	return e.value, nil
}

func (x *MinStack[T]) Peek() T {
	e, _ := x.PeekE()
	return e
}

func (x *MinStack[T]) PeekE() (T, error) {
	e, err := x.stack.PeekE()
	if err != nil {
		var zero T
		return zero, err
	}
	return e.value, nil
}

func (x *MinStack[T]) IsEmpty() bool {
	return x.stack.IsEmpty()
}

func (x *MinStack[T]) IsNotEmpty() bool {
	return x.stack.IsNotEmpty()
}

func (x *MinStack[T]) Size() int {
	return x.stack.Size()
}

func (x *MinStack[T]) Clear()  {
	x.stack.Clear()
}

func (x *MinStack[T]) String() string {
	return x.stack.String()
}

func (x *MinStack[T]) GetMin() T {
	e, _ := x.GetMinE()
	return e
}

func (x *MinStack[T]) GetMinE() (T, error) {
	e, err := x.stack.PeekE()
	if err != nil {
		var zero T
		return zero, err
	}
	return e.min, nil
}

// ------------------------------------------------ ---------------------------------------------------------------------

type MinNode[T any] struct {

	// 栈的当前位置的值
	value T

	// 栈底到当前位置的最小值
	min T
}

// ------------------------------------------------ ---------------------------------------------------------------------
