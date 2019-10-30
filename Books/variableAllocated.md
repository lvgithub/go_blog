是在堆还是堆栈上分配了变量?传递指针真的比传值效率高吗？

更多内容，欢迎关注我的[Github]([https://github.com/lvgithub/go_blog)

## 介绍
正常开发业务的角度上，其实你是没必要了解变量是分配在堆上还是栈上，完全可以把它当做黑盒。但是如果想成为一名高手，那了解这些，对未来发展还是很有帮助的。

确切的说，变量的存储位置确实对程序的性能是有影响的。Go语言中， GO编译器会把该函数的本地变量分理赔在函数的栈内存中，但是如果变化一起无法判证明本地变量在函数返回后，其他地方多没有对它引用的时候，就会把该变量分配到堆中，以避免空指针异常。或者如果该变量很大，那它也会被分配堆上。到这里也衍生出一个思考:
* 函数传递指针还是传值？
* 两种选择的本质区别是什么？
* 哪种方式的性能更高呢？

## 分析

要找到区别，那肯定需要下功夫，那就从 Golang 的实现机制中来分析吧。首先，在Golang 中有一个很重要的概念那就是 逃逸分析（Escape analysis），所谓的逃逸分析指由编译器决定内存分配的位置。
*  分配在 栈中，则函数执行结束可自动将内存回收
*  分配在 堆中，则函数执行结束可交给GC（垃圾回收）处理

最终程序的执行效率和这个两种分配规则是有这重要关联的，而传值和传指针的主要区别在于底层值是否需要拷贝,表面上看传指针不涉及值拷贝，效率肯定更高。但是实际情况是指传针会涉及到变量逃逸到堆上，而且会增加GC的负担，所以本文我们要做的内容就是进行 逃逸分析 ,按照惯例先上结论。
* 栈上分配内存比在堆中分配内存有更高的效率
* 栈上分配的内存不需要GC处理,函数执行后自动回收
* 堆上分配的内存使用完毕会交给GC处理
* 发生逃逸时，会把栈上的申请的内存移动到堆上
* 指针可以减少底层值的拷贝，可以提高效率，但是会产生逃逸，但是如果拷贝的数据量小，逃逸造成的负担（堆内存分配+GC回收)会降低效率
* 因此选择值传递还是指针传递，变量的大小是一个很重要的分析指标

每种方式都有各自的优缺点，栈上的值，减少了 GC 的压力,但是要维护多个副本，堆上的指针，会增加 GC 的压力，但只需维护一个值。因此选择哪种方式，依据自己的业务情况参考这个标准进行选择。

## 先上一段代码分析下
```golang
// escape.go
package main

type person struct {
	name string
	age  int
}

func main() {
	makePerson(32, "艾玛·斯通")
	showPerson(33, "杨幂")
}

func makePerson(age int, name string) *person {
	maliya := person{name, age}
	return &maliya
}

func showPerson(age int, name string) person {
	yangmi := person{name, age}
	return yangmi
}
```
运行如下命令,进行逃逸分析
``` shell
go build -gcflags="-m -m -l" escape.go
```
输出结果：
```
Escape/Escape.go:15:9: &maliya escapes to heap
Escape/Escape.go:15:9:  from ~r2 (return) at Escape/Escape.go:15:2
Escape/Escape.go:14:2: moved to heap: maliya
Escape/Escape.go:13:40: leaking param: name to result ~r2 level=-1
Escape/Escape.go:13:40:         from person literal (struct literal element) at Escape/Escape.go:14:18
Escape/Escape.go:13:40:         from maliya (assigned) at Escape/Escape.go:14:9
Escape/Escape.go:13:40:         from &maliya (address-of) at Escape/Escape.go:15:9
Escape/Escape.go:13:40:         from ~r2 (return) at Escape/Escape.go:15:2
Escape/Escape.go:18:39: leaking param: name to result ~r2 level=0
Escape/Escape.go:18:39:         from person literal (struct literal element) at Escape/Escape.go:19:18
Escape/Escape.go:18:39:         from yangmi (assigned) at Escape/Escape.go:19:9
Escape/Escape.go:18:39:         from ~r2 (return) at Escape/Escape.go:20:2
```
从结果中我们看到变量 &maliya 发生了逃逸,变量 yangmi 没有逃逸
```
&maliya escapes to heap from ~r2 (return) at Escape/Escape.go:15:2
moved to heap: maliya
```
所以 makePerson 返回的是指针类型，发生了逃逸，而showPerson 返回的是值类型没有逃逸。

关于变量逃逸的情况还有很多，网上有很多分析的文章，就不一一举例了，直接给出结论:
* 共享了栈上的一个值时，它就会逃逸
* 栈空间不足逃逸（比如创建一个超大的slice,超过栈空间）
* 动态类型逃逸，函数参数为interface类型（典型的fmt.Println方法）
* 闭包引用对象逃逸，其实本质还是共享了栈上的值