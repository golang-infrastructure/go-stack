package stacks

import (
	"fmt"
	"strings"
)

// LinkedStack 基于链表实现的栈
type LinkedStack[T any] struct {

	// 始终指向栈顶的节点
	top *Node[T]

	// 栈中元素的个数，为了实现O(1)计算栈大小
	elementCount int
}

var _ Stack[any] = &LinkedStack[any]{}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (x *LinkedStack[T]) Push(values ...T) {
	for _, value := range values {
		node := &Node[T]{
			value: value,
		}
		// 此时栈为空
		if x.top == nil {
			// 直接把栈顶指针指向新的节点就可以了
			x.top = node
		} else {
			// 如果栈不为空，将新的节点的next指向栈顶
			node.next = x.top
			// 同时把栈顶指针往上移动，指向这个新的节点
			x.top = node
		}
		// 栈中元素个数加1
		x.elementCount++
	}
}

func (x *LinkedStack[T]) Pop() T {
	element, _ := x.PopE()
	return element
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

	x.elementCount--

	return value, nil
}

func (x *LinkedStack[T]) Peek() T {
	element, _ := x.PeekE()
	return element
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
	return x.elementCount
}

func (x *LinkedStack[T]) Clear() {
	x.top = nil
	x.elementCount = 0
}

func (x *LinkedStack[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	currentNode := x.top
	for currentNode != nil {
		sb.WriteString(fmt.Sprintf("%#v", currentNode.value))
		currentNode = currentNode.next
	}
	sb.WriteString("]")
	return sb.String()
}

// ---------------------------------------------------------------------------------------------------------------------

// Node 用于持有栈中元素的节点
type Node[T any] struct {

	// 持有的节点
	value T

	// 指向往栈底的下一个节点的，如果为空的话表示当前节点就是栈底了
	next *Node[T]
}
