package stack

type LinkedStack[T any] struct {
	top *Node[T]
	ln  int
}

var _ Stack[any] = &LinkedStack[any]{}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (x *LinkedStack[T]) Push(values ...T) error {
	for _, value := range values {
		node := &Node[T]{
			value: value,
		}
		// 此时栈为空
		if x.top == nil {
			x.top = node
		} else {
			// 将新的节点指向栈顶，同时把栈顶指针往上移动
			node.next = x.top
			x.top = node
		}
		x.ln++
	}
	return nil
}

func (x *LinkedStack[T]) Pop() T {
	e, _ := x.PopE()
	return e
}

func (x *LinkedStack[T]) PopE() (T, error) {
	if x.top == nil {
		var zero T
		return zero, ErrStackEmpty
	}
	top := x.top
	x.top = x.top.next
	value := top.value
	top.next = nil
	var zero T
	top.value = zero
	x.ln--
	return value, nil
}

func (x *LinkedStack[T]) Peek() T {
	e, _ := x.PeekE()
	return e
}

func (x *LinkedStack[T]) PeekE() (T, error) {
	if x.top == nil {
		var zero T
		return zero, ErrStackEmpty
	}
	return x.top.value, nil
}

func (x *LinkedStack[T]) IsEmpty() bool {
	return x.top == nil
}

func (x *LinkedStack[T]) IsNotEmpty() bool {
	return x.top != nil
}

func (x *LinkedStack[T]) Size() int {
	return x.ln
}

func (x *LinkedStack[T]) Clear() error {
	x.top = nil
	x.ln = 0
	return nil
}

func (x *LinkedStack[T]) String() string {
	//TODO implement me
	panic("implement me")
}

// ---------------------------------------------------------------------------------------------------------------------

type Node[T any] struct {
	value T
	next  *Node[T]
}
