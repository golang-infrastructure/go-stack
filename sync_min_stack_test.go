package stack

import (
	"fmt"
	"testing"
)

func TestSyncMinStack(t *testing.T) {

	stack := NewSyncMinStack[int](func(a int, b int) int {
		return a - b
	})

	stack.Push(3)
	stack.Push(7)
	stack.Push(9)

	fmt.Println(stack.GetMin())

}
