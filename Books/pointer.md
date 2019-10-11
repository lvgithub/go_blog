深入理解 Golang 指针

Go中一切都通过值传递，也就是说，一个函数总是得到值传递的副本，总是会分配一个值的副本给函数参数。例如
* 将int值传递的是int值的副本；
* 指针传递指针的副本，而不是指针指向的数据；
* map 和 slice 值类似于指针，他们是指向底层存储数据结构的指针，复制map、slice的值，便不会复制他们指向的数据。具体原因可以查看[
深入理解 Slice](./slice.md)

验证
```
package main

import (
	"fmt"
)

type carListType map[string]string

var carList = make(carListType)

func main() {
	age := 10
	fmt.Printf("addr is:%p\n", &age) //addr is:0xc000018088
	sayAge(age)

	setAge(&age)
	fmt.Printf("after setAge, age is:%d\n", age) // after setAge, age is:30

	carList["honda"] = "civic"
	carList["bmw"] = "320li"

	fmt.Printf("carList is:%v\n", carList)              // carList is:map[bmw:320li honda:civic]
	fmt.Printf("carList value is:%p\n", carList)        // carList value is:0xc000098000
	fmt.Printf("carList addr is:%p\n", &carList)        // carList addr is:0x1173648
	setCar(carList)                                     // setCar carList addr is:0xc00008e000
	fmt.Printf("after setCar carList is:%v\n", carList) // after setCar carList is:map[bmw:520li honda:civic]
}

func sayAge(age int) {
	fmt.Printf("addr is:%p\n", &age)  //addr is:0xc000018098
	fmt.Printf("my age is:%d\n", age) // after setAge, age is:30
}

func setAge(age *int) {
	*age = 30
	fmt.Printf("age point value is:%p\n", age) //age point value is:0xc000080008
	fmt.Printf("age point addr is:%p\n", &age) //age point addr is:0xc00008a020
}

func setCar(carList carListType) {
	fmt.Printf("setCar carList value is:%p\n", carList) // setCar carList value is:0xc000094000
	fmt.Printf("setCar carList addr is:%p\n", &carList) // setCar carList addr is:0xc00008e020
	carList["bmw"] = "520li"
}
```
pointer 和 value 类型作为 receiver 有什么区别？主要在于你是否需要修改receiver，有如下几个注意事项：
* 如果你需要修改receiver，那必须是pointer；
* 因为 slice 和 map 是引用类型，因此这里有点微妙，他们以value作为 receiver 是可以修改receiver 的，但是如果要修改自身熟悉，比如slice的长度，那还是需要以pointer作为receiver；
* 如何receiver很大，例如一个很大的结构，那么 pointer receiver性能会更佳。可以参考[从内存分配策略(堆、栈)的角度分析,函数传递指针真的比传值效率高吗？](./escape.md)；
* 官方建议如果类型的某些方法具有 pointer receiver，那么其余的方法也保持一致，[使得方法集一致](https://golang.org/doc/faq#Should)；
* 对于基础类型、小型slice、map之类，除非强制要求，否则使用value receiver的将很高效和清晰

```
package main

import "fmt"

type man struct {
	name string
	age  int
}

type carList map[string]string

func main() {
	kangkang := man{"kangkang", 10}

	fmt.Printf("name:%s, age:%d\n", kangkang.name, kangkang.age) 
    // name:kangkang, age:10

	kangkang.setName()
	kangkang.setAge()

	fmt.Printf("name:%s, age:%d\n", kangkang.name, kangkang.age) 
    // name:kitty, age:10

	myCar :=carList{"honda":"red","bmw":"white"}
	myCar.addCar("benz","blue")

	fmt.Printf("carList: %v\n",myCar) 
    // carList: map[benz:blue bmw:white honda:red]
    // 虽然是value receiver ，依然添加成功了，符合预期
}

// method on pointer
func (m *man) setName() {
	m.name = "kitty"
}

// method on value
func (m man) setAge() {
	m.age = 30
}

func(m carList) addCar(brand string ,color string)  {
	m[brand]=color
}
```


