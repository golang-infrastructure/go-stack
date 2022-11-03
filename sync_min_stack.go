package stack

import "sync"

type SyncMinStack[T any] struct {
	stack *MinStack[T]
	lock  sync.RWMutex
}

var _ Stack[any] = &SyncMinStack[any]{}

func NewSyncMinStack[T any](comparator Comparator[T]) *SyncMinStack[T] {
	return &SyncMinStack[T]{
		stack: NewMinStack[T](comparator),
		lock:  sync.RWMutex{},
	}
}

func (x *SyncMinStack[T]) Push(values ...T) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.stack.Push(values...)
}

func (x *SyncMinStack[T]) Pop() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Pop()
}

func (x *SyncMinStack[T]) PopE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.PopE()
}

func (x *SyncMinStack[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.Peek()
}

func (x *SyncMinStack[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.PeekE()
}

func (x *SyncMinStack[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.IsEmpty()
}

func (x *SyncMinStack[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.IsNotEmpty()
}

func (x *SyncMinStack[T]) Size() int {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.Size()
}

func (x *SyncMinStack[T]) Clear() {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.stack.Clear()
}

func (x *SyncMinStack[T]) String() string {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.String()
}

func (x *SyncMinStack[T]) GetMin() T {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.GetMin()
}

func (x *SyncMinStack[T]) GetMinE() (T, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.stack.GetMinE()
}
