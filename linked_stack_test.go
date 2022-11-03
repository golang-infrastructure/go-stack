package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedStack_Clear(t *testing.T) {

}

func TestLinkedStack_IsEmpty(t *testing.T) {

}

func TestLinkedStack_IsNotEmpty(t *testing.T) {

}

func TestLinkedStack_Peek(t *testing.T) {

}

func TestLinkedStack_PeekE(t *testing.T) {

}

func TestLinkedStack_Pop(t *testing.T) {

}

func TestLinkedStack_PopE(t *testing.T) {

}

func TestLinkedStack_Push(t *testing.T) {

}

func TestLinkedStack_Size(t *testing.T) {

}

func TestLinkedStack_String(t *testing.T) {

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
