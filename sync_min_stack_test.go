package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_Clear(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleSyncMinStack_Clear() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_IsEmpty(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
}

func ExampleSyncMinStack_IsEmpty() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_IsNotEmpty(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleSyncMinStack_IsNotEmpty() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewSyncMinStack[*User](func(a, b *User) int { return 0 })
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleSyncMinStack_Peek() {
	type User struct {
	}
	stack := NewSyncMinStack[*User](func(a, b *User) int { return 0 })
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_PeekE(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncMinStack_PeekE() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
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

func TestSyncMinStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewSyncMinStack[*User](func(a, b *User) int { return 0 })
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleSyncMinStack_Pop() {
	type User struct {
	}
	stack := NewSyncMinStack[*User](func(a, b *User) int { return 0 })
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_PopE(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncMinStack_PopE() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
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

func TestSyncMinStack_Push(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
}

func ExampleSyncMinStack_Push() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_Size(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleSyncMinStack_Size() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_String(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, "[&stack.MinNode[int]{value:1, min:1}]", stack.String())
}

func ExampleSyncMinStack_String() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [&stack.MinNode[int]{value:1, min:1}]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewSyncMinStack(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	assert.NotNil(t, stack)
}

func ExampleNewSyncMinStack() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_GetMin(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	assert.Equal(t, 7, stack.GetMin())
}

func ExampleSyncMinStack_GetMin() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	fmt.Println(stack.GetMin())
	// Output:
	// 7
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMinStack_GetMinE(t *testing.T) {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })

	_, err := stack.GetMinE()
	assert.ErrorIs(t, err, ErrStackEmpty)

	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	element, err := stack.GetMinE()
	assert.Nil(t, err)
	assert.Equal(t, 7, element)
}

func ExampleSyncMinStack_GetMinE() {
	stack := NewSyncMinStack[int](func(a, b int) int { return a - b })

	_, err := stack.GetMinE()
	if errors.Is(err, ErrStackEmpty) {
		fmt.Println("stack empty!")
	}

	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	element, err := stack.GetMinE()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(element)
	// Output:
	// stack empty!
	// 7
}

// ------------------------------------------------ ---------------------------------------------------------------------
