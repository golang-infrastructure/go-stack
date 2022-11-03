package stack

type MaxStack[T any] struct {
	comparator Comparator[T]
	stack      Stack[*MaxNode[T]]
}

var _ Stack[any] = &MaxStack[any]{}

func NewMaxStack[T any](comparator Comparator[T]) *MaxStack[T] {
	return &MaxStack[T]{
		comparator: comparator,
		stack:      NewArrayStack[*MaxNode[T]](),
	}
}

func (x *MaxStack[T]) Push(values ...T) {
	for _, value := range values {
		node := &MaxNode[T]{
			value: value,
		}
		topNode, err := x.stack.PeekE()
		if err != nil {
			// 栈为空，认为最大值是自己
			node.max = value
		} else if x.comparator(topNode.max, value) > 0 {
			// 当前栈顶的元素比较小
			node.max = topNode.max
		} else {
			node.max = value
		}
		x.stack.Push(node)
	}
}

func (x *MaxStack[T]) Pop() T {
	e, _ := x.PopE()
	return e
}

func (x *MaxStack[T]) PopE() (T, error) {
	e, err := x.stack.PopE()
	if err != nil {
		var zero T
		return zero , err
	}
	return e.value, nil
}

func (x *MaxStack[T]) Peek() T {
	e, _ := x.PeekE()
	return e
}

func (x *MaxStack[T]) PeekE() (T, error) {
	e, err := x.stack.PeekE()
	if err != nil {
		var zero T
		return zero, err
	}
	return e.value, nil
}

func (x *MaxStack[T]) IsEmpty() bool {
	return x.stack.IsEmpty()
}

func (x *MaxStack[T]) IsNotEmpty() bool {
	return x.stack.IsNotEmpty()
}

func (x *MaxStack[T]) Size() int {
	return x.stack.Size()
}

func (x *MaxStack[T]) Clear()  {
	x.stack.Clear()
}

func (x *MaxStack[T]) String() string {
	return x.stack.String()
}

func (x *MaxStack[T]) GetMax() T {
	e, _ := x.GetMaxE()
	return e
}

func (x *MaxStack[T]) GetMaxE() (T, error) {
	e, err := x.stack.PeekE()
	if err != nil {
		var zero T
		return zero, err
	}
	return e.max, nil
}

// ------------------------------------------------ ---------------------------------------------------------------------

type MaxNode[T any] struct {

	// 栈的当前位置的值
	value T

	// 栈底到当前位置的最大值
	max T
}

// ------------------------------------------------ ---------------------------------------------------------------------
