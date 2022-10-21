package stack

type LinkedStack[T any] struct {
	top *Node[T]
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
	}
	return nil
}

func (x *LinkedStack[T]) Pop() T {
	e, _ := x.PopE()
	return e
}

func (x *LinkedStack[T]) PopE() (T, error) {
	if x.top == nil {
		return nil, ErrStackEmpty
	}
	top := x.top
	x.top = x.top.next
	top.next = nil
	top.value = nil
	value := top.value
	return value, nil
}

func (x *LinkedStack[T]) Peek() T {
	e, _ := x.PeekE()
	return e
}

func (x *LinkedStack[T]) PeekE() (T, error) {
	if x.top == nil {
		return nil, ErrStackEmpty
	}
	return x.top.value, nil
}

func (x *LinkedStack[T]) IsEmpty() bool {
	return x.top == nil
}

func (x *LinkedStack[T]) IsNotEmpty() bool {
	return x.top != nil
}

func (x *LinkedStack[T]) Len() int {
	count := 0
	node := x.top
	for node != nil {
		count++
		node = node.next
	}
	return count
}

func (x *LinkedStack[T]) Clear() error {
	x.top = nil
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
