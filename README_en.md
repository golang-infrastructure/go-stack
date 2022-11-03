# Go Stack



[中文文档](./README.md)

[English Document](./README_en.md)



# 1. What is this? Why choose it? What are the advantages?
This project is a Go language implementation version of various types of stacks.

- It supports generics, makes the API more convenient to use, and avoids type forced conversion
- A thread safe version is provided for each type of stack. When using it, you can focus on the business and don't need to think about locks anymore
- The supported stack types are more abundant. If there are more interesting stacks, please mention the issues call me implementation

# 2. Installation

```bash
go get github.com/CC11001100/go-stack
```
# 3. Currently implemented stack

- ArrayStack[T any]
- SyncArrayStack[T any]
- LinkedStack[T any]
- SyncLinkedStack[T any]
- MinStack[T any]
- SyncMinStack[T any]
- MaxStack[T any]
- SyncMaxStack [T any]
  
## Support Stack: 

|Struct name | Thread safety | Blocking | Special features|
| ---------------------- | :--------: | :----: | ------------------ |
| ArrayStack[T any]      |      ×      |   ×    |                    |
| SyncArrayStack[T any]  |     √      |    ×    |                    |
| LinkedStack[T any]     |      ×      |   ×    |                    |
| SyncLinkedStack[T any] |     √      |    ×    |                    |
| MinStack[T any]        |      ×      |   ×    |  O (1) Get the minimum value in the stack|
| SyncMinStack[T any]    |     √      |    ×    |  O (1) Get the minimum value in the stack|
| MaxStack[T any]        |      ×      |   ×    |  O (1) Get the maximum value in the stack|
| SyncMaxStack [T any]   |     √      |    ×    |  O (1) Get the maximum value in the stack|
# 4. Interface: Stack
  The interface 'Stack [Any]' defines some APIs that all stacks have.

##  Stack

```
Push(values ...T)
```
Example:
```go
stack := NewArrayStack[int]()
stack.Push(1)
```
### Take Top of Stack Elements

```
Pop() T
PopE() (T, error)
```
Pop Example:
```go
type User struct {
}
stack := NewArrayStack[*User]()
fmt.Println(stack.Pop())
u := &User{}
stack.Push(u)
fmt.Println(stack.Pop())
// Output:
// <nil>
// &{}
```
PopE Example:
```go
stack := NewArrayStack[int]()
element, err := stack.PopE()
if errors. Is(err, ErrStackEmpty) {
fmt. Println("stack empty!")
}
stack.Push(1)
element, err = stack.PopE()
if err !=  nil {
fmt.Println(err.Error())
return
}
fmt.Println(element)
// Output:
// stack empty!
// 1
```
### View Top of Stack Elements

```
Peek() T
PeekE() (T, error)
```
Peek Example:
```go
type User struct {
}
stack := NewArrayStack[*User]()
fmt.Println(stack.Peek())
u := &User{}
stack.Push(u)
fmt.Println(stack.Peek())
// Output:
// <nil>
// &{}
```
PeekE Example:
```go
stack := NewArrayStack[int]()
element, err := stack.PeekE()
if errors. Is(err, ErrStackEmpty) {
fmt. Println("stack empty!")
}
stack.Push(1)
element, err = stack.PeekE()
if err !=  nil {
fmt.Println(err.Error())
return
}
fmt.Println(element)
// Output:
// stack empty!
// 1
```
### Determine whether the stack is empty?

```
IsEmpty() bool
IsNotEmpty() bool
```
IsEmpty Example:
```go
stack := NewArrayStack[int]()
fmt.Println(stack.IsEmpty())
stack.Push(1)
fmt.Println(stack.IsEmpty())
// Output:
// true
// false
```
IsNotEmpty Example:
```go
stack := NewArrayStack[int]()
fmt.Println(stack.IsNotEmpty())
stack.Push(1)
fmt.Println(stack.IsNotEmpty())
// Output:
// false
// true
```
### Number of elements in the stack?

```
Size() int
```
Example:
```go
stack := NewArrayStack[int]()
stack.Push(1)
fmt.Println(stack.Size())
// Output:
// 1
```
### reset stack, clear all element!

```
Clear() 
```
Example:
```go
stack := NewArrayStack[int]()
stack.Push(1)
fmt.Println(stack.Size())
stack.Clear()
fmt.Println(stack.Size())
// Output:
// 1
// 0
```
### Stack to String
Convert the stack to String, which is generally used to visualize all elements in the stack, such as debug.
Example:
```go
stack := NewArrayStack[int]()
stack.Push(1)
fmt.Println(stack.String())
// Output:
// [1]
```
# ArrayStack

Stack based on array implementation.
Example:

```go
stack := NewArrayStack[int]()
fmt.Println(stack.String())
// Output:
// []
```
# LinkedStack
Stack based on linked list.
Example：

```go
stack := NewLinkedStack[int]()
fmt.Println(stack.String())
// Output:
// []
```
# 5. Maximum stack & minimum stack

## MinStack & SyncMinStack
Compared with the Stack interface, two methods are added for O (1) to obtain the minimum value of all elements in the stack:
```go
GetMin() T
GetMinE() (T, error)
```
GetMin Example:
```go
stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
_,  err := stack.GetMinE()
assert. ErrorIs(t, err, ErrStackEmpty)
stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMinE()
assert. Nil(t, err)
assert. Equal(t, 7, element)
```
GetMinE Example:
```go
stack := NewSyncMinStack[int](func(a, b int) int { return a - b })
_,  err := stack.GetMinE()
if errors. Is(err, ErrStackEmpty) {
fmt. Println("stack empty!")
}
stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMinE()
if err !=  nil {
fmt.Println(err.Error())
return
}
fmt.Println(element)
// Output:
// stack empty!
// 7
```
## MaxStack & SyncMaxStack
Compared with the Stack interface, two methods are added for O (1) to obtain the minimum value of all elements in the stack:
```go
GetMax() T 
GetMaxE() (T, error) 
```
GetMax Example:
```go
stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
_,  err := stack.GetMaxE()
assert. ErrorIs(t, err, ErrStackEmpty)
stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMaxE()
assert. Nil(t, err)
assert. Equal(t, 10, element)
```
GetMaxE Example:
```go
stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })
_,  err := stack.GetMaxE()
if errors. Is(err, ErrStackEmpty) {
fmt. Println("stack empty!")
}
stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMaxE()
if err !=  nil {
fmt.Println(err.Error())
return
}
fmt.Println(element)
// Output:
// stack empty!
// 10
```
# TODO
Implement the blocking stack, because only multiple coroutine operations