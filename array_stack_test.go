package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack_Clear(t *testing.T) {
	stack := NewArrayStack[int]()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Size())

}

func TestArrayStack_IsEmpty(t *testing.T) {

}

func TestArrayStack_IsNotEmpty(t *testing.T) {

}

func TestArrayStack_Size(t *testing.T) {

}

func TestArrayStack_Peek(t *testing.T) {

}

func TestArrayStack_PeekE(t *testing.T) {

}

func TestArrayStack_Pop(t *testing.T) {

}

func TestArrayStack_PopE(t *testing.T) {

}

func TestArrayStack_Push(t *testing.T) {

}

func TestArrayStack_String(t *testing.T) {

}

func TestNewArrayStack(t *testing.T) {

}
