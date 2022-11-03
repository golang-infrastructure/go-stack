package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_Clear(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleSyncArrayStack_Clear() {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_IsEmpty(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
}

func ExampleSyncArrayStack_IsEmpty() {
	stack := NewSyncArrayStack[int]()
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_IsNotEmpty(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleSyncArrayStack_IsNotEmpty() {
	stack := NewSyncArrayStack[int]()
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewSyncArrayStack[*User]()
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleSyncArrayStack_Peek() {
	type User struct {
	}
	stack := NewSyncArrayStack[*User]()
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_PeekE(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncArrayStack_PeekE() {
	stack := NewSyncArrayStack[int]()
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

func TestSyncArrayStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewSyncArrayStack[*User]()
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleSyncArrayStack_Pop() {
	type User struct {
	}
	stack := NewSyncArrayStack[*User]()
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_PopE(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncArrayStack_PopE() {
	stack := NewSyncArrayStack[int]()
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

func TestSyncArrayStack_Push(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
}

func ExampleSyncArrayStack_Push() {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_Size(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleSyncArrayStack_Size() {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncArrayStack_String(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	assert.Equal(t, "[1]", stack.String())
}

func ExampleSyncArrayStack_String() {
	stack := NewSyncArrayStack[int]()
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [1]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewSyncArrayStack(t *testing.T) {
	stack := NewSyncArrayStack[int]()
	assert.NotNil(t, stack)
}

func ExampleNewSyncArrayStack() {
	stack := NewSyncArrayStack[int]()
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------
