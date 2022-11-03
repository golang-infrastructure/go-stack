package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_Clear(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	stack.Clear()
	assert.Equal(t, 0, stack.Size())
}

func ExampleSyncMaxStack_Clear() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Clear()
	fmt.Println(stack.Size())
	// Output:
	// 1
	// 0
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_IsEmpty(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	assert.Equal(t, true, stack.IsEmpty())
	stack.Push(1)
	assert.Equal(t, false, stack.IsEmpty())
}

func ExampleSyncMaxStack_IsEmpty() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	// Output:
	// true
	// false
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_IsNotEmpty(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	assert.Equal(t, false, stack.IsNotEmpty())
	stack.Push(1)
	assert.Equal(t, true, stack.IsNotEmpty())
}

func ExampleSyncMaxStack_IsNotEmpty() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.IsNotEmpty())
	stack.Push(1)
	fmt.Println(stack.IsNotEmpty())
	// Output:
	// false
	// true
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_Peek(t *testing.T) {
	type User struct {
	}
	stack := NewSyncMaxStack[*User](func(a, b *User) int { return 0 })
	assert.Nil(t, stack.Peek())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Peek())
}

func ExampleSyncMaxStack_Peek() {
	type User struct {
	}
	stack := NewSyncMaxStack[*User](func(a, b *User) int { return 0 })
	fmt.Println(stack.Peek())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Peek())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_PeekE(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	element, err := stack.PeekE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncMaxStack_PeekE() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
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

func TestSyncMaxStack_Pop(t *testing.T) {
	type User struct {
	}
	stack := NewSyncMaxStack[*User](func(a, b *User) int { return 0 })
	assert.Nil(t, stack.Pop())
	u := &User{}
	stack.Push(u)
	assert.Equal(t, u, stack.Pop())
}

func ExampleSyncMaxStack_Pop() {
	type User struct {
	}
	stack := NewSyncMaxStack[*User](func(a, b *User) int { return 0 })
	fmt.Println(stack.Pop())
	u := &User{}
	stack.Push(u)
	fmt.Println(stack.Pop())
	// Output:
	// <nil>
	// &{}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_PopE(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	element, err := stack.PopE()
	assert.Equal(t, 0, element)
	assert.ErrorIs(t, ErrStackEmpty, err)

	stack.Push(1)
	element, err = stack.PopE()
	assert.Nil(t, err)
	assert.Equal(t, 1, element)
}

func ExampleSyncMaxStack_PopE() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
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

func TestSyncMaxStack_Push(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
}

func ExampleSyncMaxStack_Push() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	// Output:
	//
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_Size(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
}

func ExampleSyncMaxStack_Size() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.Size())
	// Output:
	// 1
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_String(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	assert.Equal(t, "[&stack.MaxNode[int]{value:1, max:1}]", stack.String())
}

func ExampleSyncMaxStack_String() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(1)
	fmt.Println(stack.String())
	// Output:
	// [&stack.MaxNode[int]{value:1, max:1}]
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestNewSyncMaxStack(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	assert.NotNil(t, stack)
}

func ExampleNewSyncMaxStack() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	fmt.Println(stack.String())
	// Output:
	// []
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_GetMax(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	assert.Equal(t, 10, stack.GetMax())
}

func ExampleSyncMaxStack_GetMax() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	fmt.Println(stack.GetMax())
	// Output:
	// 10
}

// ------------------------------------------------ ---------------------------------------------------------------------

func TestSyncMaxStack_GetMaxE(t *testing.T) {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })

	_, err := stack.GetMaxE()
	assert.ErrorIs(t, err, ErrStackEmpty)

	stack.Push(10)
	stack.Push(7)
	stack.Push(9)
	element, err := stack.GetMaxE()
	assert.Nil(t, err)
	assert.Equal(t, 10, element)
}

func ExampleSyncMaxStack_GetMaxE() {
	stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })

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
