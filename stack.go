package stack

// Stack 定义栈的API
type Stack[T any] interface {

	// Push 往栈中加入元素，可以一次增加多个
	Push(values ...T) error

	// Pop 弹出栈顶元素，如果栈为空的话返回栈元素T对应的零值
	Pop() T
	// PopE 弹出栈顶元素，如果栈为空的话则返回 ErrStackEmpty 错误
	PopE() (T, error)

	// Peek 访问但不弹出栈顶元素，如果栈为空的话则返回栈元素T对应的零值
	Peek() T
	// PeekE 访问但不弹出栈顶元素，如果栈为空的话则返回 ErrStackEmpty 错误
	PeekE() (T, error)

	// IsEmpty 栈是否为空
	IsEmpty() bool
	// IsNotEmpty 栈是否不为空
	IsNotEmpty() bool

	// Size 返回栈中元素的个数
	Size() int

	// Clear 清空栈
	Clear() error

	// String 把栈打印为字符串，一般debug之类的可能会用得到
	String() string
}

// NewStack 默认是用数组实现的栈
func NewStack[T any]() Stack[T] {
	return NewArrayStack[T]()
}
