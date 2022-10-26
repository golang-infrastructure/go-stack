package stack

//import "context"
//
//type BlockingStackImpl[T any] struct {
//	stack chan T
//}
//
//var _ BlockingStack[any] = &BlockingStackImpl[any]{}
//
//func NewBlockingStack[T any](capacity int) *BlockingStackImpl[T] {
//	return &BlockingStackImpl[T]{
//		stack: make(chan T, capacity),
//	}
//}
//
//func (x *BlockingStackImpl[T]) Push(values ...T) error {
//}
//
//func (x *BlockingStackImpl[T]) Pop() T {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) PopE() (T, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) Peek() T {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) PeekE() (T, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) IsEmpty() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) IsNotEmpty() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) Len() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) Clear() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) String() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) BPush(ctx context.Context, values ...T) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) BPop(ctx context.Context) T {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *BlockingStackImpl[T]) BPopE(ctx context.Context) (T, error) {
//	select {
//	case value := <-x.stack:
//		return value, nil
//	case <-ctx.Done():
//		return nil, ErrContextTimeout
//	}
//}
//
//func (x *BlockingStackImpl[T]) BPeek(ctx context.Context) T {
//	e, _ := x.BPeekE(ctx)
//	return e
//}
//
//func (x *BlockingStackImpl[T]) BPeekE(ctx context.Context) (T, error) {
//	select {
//	case value := <-x.stack:
//		return value, nil
//	case <-ctx.Done():
//		return nil, ErrContextTimeout
//	}
//}
