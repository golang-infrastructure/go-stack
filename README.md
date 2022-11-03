# Go Stack 

# 一、这是什么？为什么使用它？优势是啥？

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

| Struct名称             | 线程安全性 | 阻塞性 | 特有特性 |
| ---------------------- | :--------: | :----: | -------- |
| ArrayStack[T any]      |     ×      |   ×    |          |
| SyncArrayStack[T any]  |     √      |   ×    |          |
| LinkedStack[T any]     |     ×      |   ×    |          |
| SyncLinkedStack[T any] |     √      |   ×    |          |
| MinStack[T any]        |     ×      |   ×    |          |
| SyncMinStack[T any]    |     √      |   ×    |          |
| MaxStack[T any]        |     ×      |   ×    |          |
| SyncMaxStack [T any]   |     √      |   ×    |          |



# 四、Interface: Stack 

接口定义的接口定义



## 入栈

```
Push(values ...T)
```



## 出栈

```
Pop() T
PopE() (T, error)
```



## 查看栈顶元素

```
Peek() T
PeekE() (T, error)
```



## 判断栈是否空

```
IsEmpty() bool
IsNotEmpty() bool
```



## 栈中元素个数

```
Size() int
```



## 清空栈

```
Clear() 
```



## ArrayStack

数组实现的栈适合频繁入栈出栈的操作

### 数组缩容

可能会存在的坑：类似于GO内置的map，默认情况下ArrayStack底层的数组也是只增长不缩容的，如果你的栈中的元素的个数具有先是很多，然后一直保持在一个较少的样子，建议在栈相对平稳之后调用TrimArray方法来让底层的数组释放掉短时间内不会再使用的空间。

### 数组扩容到指定长度

允许对底层数组进行一些微操。



## SyncArrayStack

## LinkedStack

## SyncLinkedStack



# 五、最大栈 & 最小栈

## MinStack

## SyncMinStack

## MaxStack

SyncMaxStack 



# 六、阻塞栈

因为仅当多个协程操作同一个栈时才需要考虑阻塞的情况，所以阻塞栈都是线程安全的。





# TODO 

TODO 2022-10-22 02:07:40 Code Review  
TODO 2022-10-22 02:04:55 测试  
TODO 2022-10-22 02:03:14 编写文档   





