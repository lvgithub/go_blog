Golang 中 new 和 make的区别

Golang 内置两个函数，new 和 make ,其作用都是用来分配内存的。这么说可能会造成混淆，但其实规则很简单，其实他们适用于不同的类型。
* make 函数只用于给 slice、map、channel进行初始化；
* new 函数参数为 一个类型不是一个值，用该类型的零值初始化一块空间，并返回该空间的指针。

我们看一下下面的案例：
```
package main

import (
	"fmt"
)

func main() {
	var i *int

	*i=10
	fmt.Println(*i)

}
```
运行报错，panic: runtime error: invalid memory address or nil pointer dereference。错误很明显提示无效的内存地址或者空指针错误。  
如何解决呢，很简单，我们只需要使用new来进行内存分配。
```
package main

import (
	"fmt"
)

func main() {
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)

}
```
再次运行，完美输出结果 10 。  
我们再通过源码看看new 函数
```
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
```
只接收一个参数（type），并且返回了一个初始化为零值的内存指针。

再看看 make 函数
```
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.
func make(t Type, size ...IntegerType) Type
```
* 只为 slice, map, or chan (only)分配内存和初始化；
* 和 new 函数一样的是第一个参数是一个 type 不是 value；
* 和 new 函数不同的是make 返回的不是指针，返回的是参数中传递进来的类型本身；

你肯定会困惑，为什么make 不返回指针呢，因为slice, map, or chan本身就是引用类型。具体可以查看文章[深入理解 Slice](https://github.com/lvgithub/go_blog/blob/master/Books/slice.md)

更多内容，欢迎关注我的[Github]([https://github.com/lvgithub/go_blog)。