package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_Clear(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleMaxStack_Clear() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_IsEmpty(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
}

func ExampleMaxStack_IsEmpty() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_IsNotEmpty(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleMaxStack_IsNotEmpty() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewMaxStack[*User](func(a, b *User) int { return 0 })
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleMaxStack_Peek() {
	type User struct {
	}
	stack := NewMaxStack[*User](func(a, b *User) int { return 0 })
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_PeekE(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleMaxStack_PeekE() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
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

func TestMaxStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewMaxStack[*User](func(a, b *User) int { return 0 })
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleMaxStack_Pop() {
	type User struct {
	}
	stack := NewMaxStack[*User](func(a, b *User) int { return 0 })
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_PopE(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleMaxStack_PopE() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
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

func TestMaxStack_Push(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
}

func ExampleMaxStack_Push() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_Size(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleMaxStack_Size() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_String(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, "[&stack.MaxNode[int]{value:1, max:1}]", stack.String())
}

func ExampleMaxStack_String() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [&stack.MaxNode[int]{value:1, max:1}]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewMaxStack(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	assert.NotNil(t, stack)
}

func ExampleNewMaxStack() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_GetMax(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	assert.Equal(t, 10, stack.GetMax())
}

func ExampleMaxStack_GetMax() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	fmt.Println(stack.GetMax())
	// Output:
	// 10
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestMaxStack_GetMaxE(t *testing.T) {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })

	_, err := stack.GetMaxE()
	assert.ErrorIs(t, err, ErrStackEmpty)

	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	element, err := stack.GetMaxE()
	assert.Nil(t, err)
	assert.Equal(t, 10, element)
}

func ExampleMaxStack_GetMaxE() {
	stack := NewMaxStack[int](func(a, b int) int { return a - b })

	_, err := stack.GetMaxE()
	if errors.Is(err, ErrStackEmpty) {
		fmt.Println("stack empty!")
	}

	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	element, err := stack.GetMaxE()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(element)
	// Output:
	// stack empty!
	// 10
}

// ------------------------------------------------ ---------------------------------------------------------------------
