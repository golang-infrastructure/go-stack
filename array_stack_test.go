package stacks

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Clear(t *testing.T) {
	stack := NewArrayStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleArrayStack_Clear() {
	stack := NewArrayStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_IsEmpty(t *testing.T) {
	stack := NewArrayStack[int]()
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
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
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleArrayStack_IsNotEmpty() {
	stack := NewArrayStack[int]()
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewArrayStack[*User]()
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleArrayStack_Peek() {
	type User struct {
	}
	stack := NewArrayStack[*User]()
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_PeekE(t *testing.T) {
	stack := NewArrayStack[int]()
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleArrayStack_PeekE() {
	stack := NewArrayStack[int]()
	element, err := stack.PeekE()
	if errors.Is(err, ErrStackEmpty) {
		fmt.Println("stack empty!")
	}

	stack.Push(1)
	element, err = stack.PeekE()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(element)
	// Output:
	// stack empty!
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewArrayStack[*User]()
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleArrayStack_Pop() {
	type User struct {
	}
	stack := NewArrayStack[*User]()
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_PopE(t *testing.T) {
	stack := NewArrayStack[int]()
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleArrayStack_PopE() {
	stack := NewArrayStack[int]()
	element, err := stack.PopE()
	if errors.Is(err, ErrStackEmpty) {
		fmt.Println("stack empty!")
	}

	stack.Push(1)
	element, err = stack.PopE()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(element)
	// Output:
	// stack empty!
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack[int]()
	stack.Push(1)
}

func ExampleArrayStack_Push() {
	stack := NewArrayStack[int]()
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_Size(t *testing.T) {
	stack := NewArrayStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleArrayStack_Size() {
	stack := NewArrayStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestArrayStack_String(t *testing.T) {
	stack := NewArrayStack[int]()
	stack.Push(1)
	assert.Equal(t, "[1]", stack.String())
}

func ExampleArrayStack_String() {
	stack := NewArrayStack[int]()
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [1]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewArrayStack(t *testing.T) {
	stack := NewArrayStack[int]()
	assert.NotNil(t, stack)
}

func ExampleNewArrayStack() {
	stack := NewArrayStack[int]()
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------
