package stack

import "sync"

type SyncMaxStack[T any] struct {
	stack *MaxStack[T]
	lock  sync.RWMutex
}

var _ Stack[any] = &SyncMaxStack[any]{}

func NewSyncMaxStack[T any](comparator Comparator[T]) *SyncMaxStack[T] {
	return &SyncMaxStack[T]{
		stack: NewMaxStack[T](comparator),
		lock:  sync.RWMutex{},
	}
}

func (x *SyncMaxStack[T]) Push(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Push(values...)
}

func (x *SyncMaxStack[T]) Pop() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Pop()
}

func (x *SyncMaxStack[T]) PopE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.PopE()
}

func (x *SyncMaxStack[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Peek()
}

func (x *SyncMaxStack[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.PeekE()
}

func (x *SyncMaxStack[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsEmpty()
}

func (x *SyncMaxStack[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsNotEmpty()
}

func (x *SyncMaxStack[T]) Len() int {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Len()
}

func (x *SyncMaxStack[T]) Clear() error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Clear()
}

func (x *SyncMaxStack[T]) String() string {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.String()
}

func (x *SyncMaxStack[T]) GetMax() T {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.GetMax()
}

func (x *SyncMaxStack[T]) GetMaxE() (T, error) {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.GetMaxE()
}
