package stack

import "sync"

type SyncArrayStack[T any] struct {
	stack *ArrayStack[T]
	lock  sync.RWMutex
}

var _ Stack[any] = &SyncArrayStack[any]{}

func NewSyncArrayStack[T any]() *SyncArrayStack[T] {
	return &SyncArrayStack[T]{
		stack: NewArrayStack[T](),
		lock:  sync.RWMutex{},
	}
}

func (x *SyncArrayStack[T]) Push(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Push(values...)
}

func (x *SyncArrayStack[T]) Pop() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Pop()
}

func (x *SyncArrayStack[T]) PopE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.PopE()
}

func (x *SyncArrayStack[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Peek()
}

func (x *SyncArrayStack[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.PeekE()
}

func (x *SyncArrayStack[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsEmpty()
}

func (x *SyncArrayStack[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsNotEmpty()
}

func (x *SyncArrayStack[T]) Len() int {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Len()
}

func (x *SyncArrayStack[T]) Clear() error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Clear()
}

func (x *SyncArrayStack[T]) String() string {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.String()
}
