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

func (x *SyncMinStack[T]) Push(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Push(values...)
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
	defer x.lock.RLock()
	return x.stack.Peek()
}

func (x *SyncMinStack[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.PeekE()
}

func (x *SyncMinStack[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsEmpty()
}

func (x *SyncMinStack[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.IsNotEmpty()
}

func (x *SyncMinStack[T]) Len() int {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.Len()
}

func (x *SyncMinStack[T]) Clear() error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.stack.Clear()
}

func (x *SyncMinStack[T]) String() string {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.String()
}

func (x *SyncMinStack[T]) GetMin() T {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.GetMin()
}

func (x *SyncMinStack[T]) GetMinE() (T, error) {
	x.lock.RLock()
	defer x.lock.RLock()
	return x.stack.GetMinE()
}