package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Clear(t *testing.T) {
	stack := NewArrayStack[int]()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleArrayStack_Clear() {
	stack := NewArrayStack[int]()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 2
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_IsEmpty(t *testing.T) {
	stack := NewArrayStack[int]()
	assert.Equal(t, stack.IsEmpty(), true)
	stack.Push(1)
	assert.Equal(t, stack.IsEmpty(), false)
}

func ExampleArrayStack_IsEmpty() {
	stack := NewArrayStack[int]()
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_IsNotEmpty(t *testing.T) {
	stack := NewArrayStack[int]()
	assert.Equal(t, stack.IsNotEmpty(), false)
	stack.Push(1)
	assert.Equal(t, stack.IsEmpty(), true)
}

func ExampleArrayStack_IsNotEmpty() {
	stack := NewArrayStack[int]()
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Size(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Peek(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_PeekE(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Pop(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_PopE(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Push(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_String(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewArrayStack(t *testing.T) {

}

// ------------------------------------------------ ---------------------------------------------------------------------
