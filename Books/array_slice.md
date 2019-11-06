通过 深入对比 Arrays 和 Slices 学习GO

## 数组定义
```
package main

import (
	"fmt"
)

func main() {
	var arr [3]int      //int 数组，长度为3 ，并且初始值都为0 
	fmt.Println(arr)    // output:[0 0 0]
}
```
因为很简单的基础知识，就不一一解释了。

## 数组使用
```
package main

import (
	"fmt"
)

func main() {
	arr := [3]int{12, 78, 50} //int 数组，长度为3 ，并且初始化
	fmt.Println(arr)          // output:[12 78 50]
}
```
上面通过简写的方式声明了数组，并且赋值，但并不是需要为所有元素分配值。如下程序：
```
package main

import (
	"fmt"
)

func main() {
	arr := [3]int{12} //int 数组，长度为3 ，并且初始化一个值
	fmt.Println(arr)          // output:[12 0 0]
}
```
上面程序中哪个 arr 是一个长度为3的数组，只给第0为初始化为12，其他值依旧被默认初始化为0。善于“偷懒”的程序员往往是高手，比如你可以忽略数组的长度，让编译器帮我们去设置，如下：
```
package main

import (
	"fmt"
)

func main() {
	arr := [...]int{12, 78, 99} //int 数组，长度为3 ，并且初始化一个值
	fmt.Println(arr)            // output:[12 0 0]
}
```
通过 *...* 编译器可以自动找到数组长度，不同长度的数组类型是属于不同类型，比如：
```
package main

func main() {
	a := [3]int{5, 78, 8}
	var b [5]int
	b = a //cannot use a (type [3]int) as type [5]int in assignment
}
```
编译器直接抛出错误： *cannot use a (type [3]int) as type [5]int in assignment*

Arrays 属于值类型，也就意味着，如果你将数组赋值给一个新的变量，对新的变量进行修改，将不会影响原有的变量。
```
package main

import "fmt"

func main() {
	a := [3]int{5, 78, 8}
	var b = a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
```
output:
```
[5 78 8]
[100 78 8]
```
修改了b,a没有被修改，可见数组的赋值是一个值拷贝。

## Slice 定义
```
package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // 基于数组 a 创建一个 slice b
	fmt.Println(b)       // [77 78 79]
}
```
上面我们创建了一个 *slice* 基于数组A，开始于数组的第一位结束于数组的第四位（不包含），然后我们看看改变数组a 会发生什么：
```
package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // 基于数组 a 创建一个 slice b
	fmt.Println(b)       // [77 78 79]

	a[1] = 100
	fmt.Println(b) // [100 78 79]
}
```
修改了a[1],结果b[1]的值也修改了，可以看到 slice 是对 数组 a 的应用， slice b 的底层存储实际是数组 b。 为什么是这样呢？可以查看文章[深入理解 Slice](https://github.com/lvgithub/go_blog/blob/master/Books/slice.md)

如果我不想基于某个数组去创建 Slice 可以吗，也是可以的：
```
package main

import (
	"fmt"
)

func main() {
	i := make([]int, 5, 5)
	fmt.Println(i)

	j := []int{}
	j = append(j, 1)
	fmt.Println(j)
}
```
上面通过两种方式创建了 Slice ，尤其注意，make 的方式可以设置 slice 的长度和容量。
```
package main

import (
	"fmt"
)

func main() {
	i := make([]int, 5, 5)
	fmt.Println(i)

	var j []int

	if j == nil {
		fmt.Printf("before append: j is nil\n")
	}
	j = append(j, 1)
	if j == nil {
		fmt.Printf("after append: j is nil\n")
	}
	fmt.Printf("i len：%d  cap:%d \n", len(i), cap(i))
}

```
output:
```
[0 0 0 0 0]
before append: j is nil
i len：5  cap:5 
```
如上程序，Go 为 slice 提供了两个内置的函数 *len() cap()*用来获取 slice 的长度和容量,*append* 来给 slice 扩充数据。不像数组一样直接通过 *length* 属性来获取长度。 还要一个细节 slice 不是值类型，因此可以通过 *j == nil* 来判断,数组是值类型如果这样会报错：
```
package main

func main() {
	i := [2]int{3, 1}
	if i == nil { //cannot convert nil to type [2]int

	}
}
```
编译报错：*cannot convert nil to type [2]int*

## 内存优化
slice 在一定的情况下，非常有利于我们优化内存，不需要开辟新的内存空间，这时候只要切片在内存中，就无法对被引用的数组进行垃圾回收，所以当我们遇到一个非常大的数组时，而我们只对其中小部分感兴趣，此时我们应该使用 *copy* 函数生成该 slice 的副本，使得原数组可以被回收,代码如下：
```
package main

import (
	"fmt"
)

func main() {
	i := []int{1, 3, 4, 56, 3}
	fmt.Println(i)

	j := make([]int, 5, 5)
	copy(j, i)
	fmt.Println(j)
}
```
## 参考
[Arrays and Slices](https://golangbot.com/arrays-and-slices/)