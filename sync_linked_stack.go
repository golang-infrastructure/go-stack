package stack

import "sync"

type SyncLinkedStack[T any] struct {
	stack *LinkedStack[T]
	lock  sync.RWMutex
}

var _ Stack[any] = &SyncLinkedStack[any]{}

func NewSyncLinkedStack[T any]() *SyncLinkedStack[T] {
	return &SyncLinkedStack[T]{
		stack: NewLinkedStack[T](),
		lock:  sync.RWMutex{},
	}
}

func (x *SyncLinkedStack[T]) Push(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Push(values...)
}

func (x *SyncLinkedStack[T]) Pop() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Pop()
}

func (x *SyncLinkedStack[T]) PopE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.PopE()
}

func (x *SyncLinkedStack[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Peek()
}

func (x *SyncLinkedStack[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.PeekE()
}

func (x *SyncLinkedStack[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsEmpty()
}

func (x *SyncLinkedStack[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsNotEmpty()
}

func (x *SyncLinkedStack[T]) Size() int {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Size()
}

func (x *SyncLinkedStack[T]) Clear() error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Clear()
}

func (x *SyncLinkedStack[T]) String() string {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.String()
}
