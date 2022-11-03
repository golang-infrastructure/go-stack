package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_Clear(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleSyncLinkedStack_Clear() {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_IsEmpty(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
}

func ExampleSyncLinkedStack_IsEmpty() {
	stack := NewSyncLinkedStack[int]()
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_IsNotEmpty(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleSyncLinkedStack_IsNotEmpty() {
	stack := NewSyncLinkedStack[int]()
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewSyncLinkedStack[*User]()
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleSyncLinkedStack_Peek() {
	type User struct {
	}
	stack := NewSyncLinkedStack[*User]()
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_PeekE(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncLinkedStack_PeekE() {
	stack := NewSyncLinkedStack[int]()
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

func TestSyncLinkedStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewSyncLinkedStack[*User]()
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleSyncLinkedStack_Pop() {
	type User struct {
	}
	stack := NewSyncLinkedStack[*User]()
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_PopE(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncLinkedStack_PopE() {
	stack := NewSyncLinkedStack[int]()
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

func TestSyncLinkedStack_Push(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
}

func ExampleSyncLinkedStack_Push() {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_Size(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleSyncLinkedStack_Size() {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncLinkedStack_String(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	assert.Equal(t, "[1]", stack.String())
}

func ExampleSyncLinkedStack_String() {
	stack := NewSyncLinkedStack[int]()
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [1]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewSyncLinkedStack(t *testing.T) {
	stack := NewSyncLinkedStack[int]()
	assert.NotNil(t, stack)
}

func ExampleNewSyncLinkedStack() {
	stack := NewSyncLinkedStack[int]()
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------
