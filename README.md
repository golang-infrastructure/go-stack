# Go Stack 

# 一、这是什么？为什么选它？优势是啥？

这个项目是各种类型的栈的Go语言实现版本。

- 支持泛型，API使用更方便，避免类型强转
- 对每种类型的栈都提供了线程安全版本，使用的时候可以专注业务不需要再考虑锁的问题了  
- 支持的栈类型更丰富，如果有更多有意思的栈，请提issues call me实现

# 二、安装 

```bash
go get github.com/CC11001100/go-stack
```

# 三、目前实现的栈

- ArrayStack[T any]
- SyncArrayStack[T any]
- LinkedStack[T any]
- SyncLinkedStack[T any]
- MinStack[T any]
- SyncMinStack[T any]
- MaxStack[T any]
- SyncMaxStack [T any]

### 一览表

| Struct名称             | 线程安全性 | 阻塞性 | 特有特性           |
| ---------------------- | :--------: | :----: | ------------------ |
| ArrayStack[T any]      |     ×      |   ×    |                    |
| SyncArrayStack[T any]  |     √      |   ×    |                    |
| LinkedStack[T any]     |     ×      |   ×    |                    |
| SyncLinkedStack[T any] |     √      |   ×    |                    |
| MinStack[T any]        |     ×      |   ×    | O(1)获取栈中最小值 |
| SyncMinStack[T any]    |     √      |   ×    | O(1)获取栈中最小值 |
| MaxStack[T any]        |     ×      |   ×    | O(1)获取栈中最大值 |
| SyncMaxStack [T any]   |     √      |   ×    | O(1)获取栈中最大值 |

# 四、Interface: Stack 

接口`Stack[T any]`定义了一些所有的栈都会有的API。

## 入栈

```
Push(values ...T)
```

Example:

```go
stack := NewArrayStack[int]()
stack.Push(1)
```

## 出栈

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
if errors.Is(err, ErrStackEmpty) {
    fmt.Println("stack empty!")
}

stack.Push(1)
element, err = stack.PopE()
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Println(element)
// Output:
// stack empty!
// 1
```

## 查看栈顶元素

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
if errors.Is(err, ErrStackEmpty) {
    fmt.Println("stack empty!")
}

stack.Push(1)
element, err = stack.PeekE()
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Println(element)
// Output:
// stack empty!
// 1
```

## 判断栈是否空

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

## 栈中元素个数

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

## 清空栈

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

## String

把栈转为String，一般用于debug之类的把栈中所有元素可视化。

Example:

```go
stack := NewArrayStack[int]()
stack.Push(1)
fmt.Println(stack.String())
// Output:
// [1]
```

# ArrayStack

基于数组实现的栈。

Example:

```go
stack := NewArrayStack[int]()
fmt.Println(stack.String())
// Output:
// []
```

# LinkedStack

基于链表实现的栈。

Example：

```go
stack := NewLinkedStack[int]()
fmt.Println(stack.String())
// Output:
// []
```

# 五、最大栈 & 最小栈

## MinStack && SyncMinStack

相较于Stack接口增加了两个方法用于O(1)获取栈中所有元素的最小值：

```go
GetMin() T
GetMinE() (T, error)
```

GetMin Example:

```go
stack := NewSyncMinStack[int](func(a, b int) int { return a - b })

_, err := stack.GetMinE()
assert.ErrorIs(t, err, ErrStackEmpty)

stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMinE()
assert.Nil(t, err)
assert.Equal(t, 7, element)
```

GetMinE Example:

```go
stack := NewSyncMinStack[int](func(a, b int) int { return a - b })

_, err := stack.GetMinE()
if errors.Is(err, ErrStackEmpty) {
    fmt.Println("stack empty!")
}

stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMinE()
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Println(element)
// Output:
// stack empty!
// 7
```

## MaxStack && SyncMaxStack 

相较于Stack接口增加了两个方法用于O(1)获取栈中所有元素的最小值：

```go
GetMax() T 
GetMaxE() (T, error) 
```

GetMax Example:

```go
stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })

_, err := stack.GetMaxE()
assert.ErrorIs(t, err, ErrStackEmpty)

stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMaxE()
assert.Nil(t, err)
assert.Equal(t, 10, element)
```

GetMaxE Example:

```go
stack := NewSyncMaxStack[int](func(a, b int) int { return a - b })

_, err := stack.GetMaxE()
if errors.Is(err, ErrStackEmpty) {
    fmt.Println("stack empty!")
}

stack.Push(10)
stack.Push(7)
stack.Push(9)
element, err := stack.GetMaxE()
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Println(element)
// Output:
// stack empty!
// 10
```

# TODO 

实现阻塞栈，因为仅当多个协程操作同一个栈时才需要考虑阻塞的情况，所以阻塞栈都是线程安全的。





