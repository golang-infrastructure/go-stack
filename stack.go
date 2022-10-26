package stack

// Stack 定义栈的API
type Stack[T any] interface {

	// Push 往栈中加入元素
	Push(values ...T) error

	Pop() T
	PopE() (T, error)

	Peek() T
	PeekE() (T, error)

	IsEmpty() bool
	IsNotEmpty() bool

	Len() int

	Clear() error

	String() string
}

// NewStack 默认是用数组实现的栈
func NewStack[T any]() Stack[T] {
	return NewArrayStack[T]()
}
