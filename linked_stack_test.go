package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_Clear(t *testing.T) {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleLinkedStack_Clear() {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_IsEmpty(t *testing.T) {
	stack := NewLinkedStack[int]()
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
}

func ExampleLinkedStack_IsEmpty() {
	stack := NewLinkedStack[int]()
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_IsNotEmpty(t *testing.T) {
	stack := NewLinkedStack[int]()
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleLinkedStack_IsNotEmpty() {
	stack := NewLinkedStack[int]()
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewLinkedStack[*User]()
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleLinkedStack_Peek() {
	type User struct {
	}
	stack := NewLinkedStack[*User]()
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_PeekE(t *testing.T) {
	stack := NewLinkedStack[int]()
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleLinkedStack_PeekE() {
	stack := NewLinkedStack[int]()
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

func TestLinkedStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewLinkedStack[*User]()
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleLinkedStack_Pop() {
	type User struct {
	}
	stack := NewLinkedStack[*User]()
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_PopE(t *testing.T) {
	stack := NewLinkedStack[int]()
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleLinkedStack_PopE() {
	stack := NewLinkedStack[int]()
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

func TestLinkedStack_Push(t *testing.T) {
	stack := NewLinkedStack[int]()
	stack.Push(1)
}

func ExampleLinkedStack_Push() {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_Size(t *testing.T) {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleLinkedStack_Size() {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestLinkedStack_String(t *testing.T) {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	assert.Equal(t, "[1]", stack.String())
}

func ExampleLinkedStack_String() {
	stack := NewLinkedStack[int]()
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [1]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewLinkedStack(t *testing.T) {
	stack := NewLinkedStack[int]()
	assert.NotNil(t, stack)
}

func ExampleNewLinkedStack() {
	stack := NewLinkedStack[int]()
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------
