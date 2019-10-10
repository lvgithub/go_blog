原来这才是 Golang Interface
## 定义
* Interface 是一个定义了方法签名的集合,用来指定对象的行为，如果对象做到了 Interface 中方法集定义的行为，那就可以说实现了 Interface。

* 这些方法可以在不同的地方被不同的对象实现，这些实现可以具有不同的行为。

* interface 的主要工作仅是提供方法名称签名,输入参数,返回类型。最终由具体的对象来实现方法，比如 struct。

* interface 初始化值为 nil

* 使用 type 关键字来申明，interface 代表类型，大括号里面定义接口的方法签名集合。
	```
	type Animal interface {
		Bark() string
		Walk() string
	}
	```
	如下，Dog 实现了 Animal 接口，所以可以用 Animal 的实例去接收 Dog的实例，必须是同时实现 Bark() 和Walk() 方法，否则都不能算实现了Animal接口。
	```
	type Dog struct {
		name string
	}

	func (dog Dog) Bark() {
		fmt.Println(dog.name + ":wan wan wan!")
	}

	func (dog Dog) Walk() {
		fmt.Println(dog.name + ":walk to park!")
	}

	func main() {
		var animal Animal

		fmt.Println("animal value is:", animal)	//animal value is: <nil>
		fmt.Printf("animal type is: %T\n", animal) //animal type is: <nil>

		animal = Dog{"旺财"}
		animal.Bark() //旺财:wan wan wan!
		animal.Walk() //旺财:walk to park!

		fmt.Println("animal value is:", animal) //animal value is: {旺财}
		fmt.Printf("animal type is: %T\n", animal) //animal type is: main.Dog
	}
	```
## nil interface
在上面的例子中，我们打印刚定义的 animal:
* value为 nil
* type 也为 nil

官方定义：Interface values with nil underlying values:
* 只声明没赋值的interface 是nil interface，value和 type 都是 nil 
* 只要赋值了，即使赋了一个值为nil类型，也不再是nil interface
```
type I interface {
	Hello()
}

type S []int

func (i S) Hello() {
	fmt.Println("hello")
}
func main() {
	var i I
	fmt.Printf("1:i Type:%T\n", i)
	fmt.Printf("2:i Value:%v\n", i)

	var s S
	if s == nil {
		fmt.Printf("3:s Value%v\n", s)
		fmt.Printf("4:s Type is %T\n", s)
	}

	i = s
	if i == nil {
		fmt.Println("5:i is nil")
	} else {
		fmt.Printf("6:i Type:%T\n", i)
		fmt.Printf("7:i Value:%v\n", i)
	}
}
```
output:
```
	1:i Type:<nil>
	2:i Value:<nil>
	3:s Value[]
	4:s Type is main.S
	6:i Type:main.S
	7:i Value:[]
```
从结果看，初始化的变量 i 是一个 nil interface,当把值为 nil 的变量 s 赋值i后,i 不再为nil interface。  
细心的同学，会发现一个细节，输出的第3行
```
3:s Value[]
```
明明，s的值是 nil,却输出的是一个[],这是由于 fmt使用反射来确定打印的内容，因为 s 的类型是slice，所以 fmt用 []来表示。

## empty interface  
Go 允许不带任何方法的 interface ,这种类型的 interface 叫 empty interface。所有类型都实现了 empty interface,因为任何一种类型至少实现了 0 个方法。  
典型的应用场景是 fmt包的Println方法，它能支持接收各种不同的类型的数据，并且输出到控制台,就是interface{}的功劳。下面我们看下案例：
```
func Print(i interface{}) {
	fmt.Println(i)
}
func main() {
	var i interface{}
	i = "hello"
	Print(i)
	i = 100
	Print(i)
	i = 1.29
	Print(i)
}
```
Print 方法的参数类型为 interface{},我们传入 string,int,float等类型它都能接收。  
虽然interface{}可以接收任何类型的参数，但是interface{}类型的 slice 是不是就可以接受任何类型的 slice。如下代码将会触发 panic 错误,
```
var dataSlice []int = foo()
var interfaceSlice []interface{} = dataSlice
// cannot use dataSlice (type []int) as type []interface { } in assignment
```
具体原因，官网 wiki(https://github.com/golang/go/wiki/InterfaceSlice) 有描述,大致含义是，导致错误是有两个原因的：
* []interface{} 并不是一个interface，它是一个slice,只是slice 中的元素是interface
* []interface{} 类型的内存大小是在编译期间就确定的(N*2),而其他切片类型的大小则为 N * sizeof(MyType),因此不发快速的将类型[]MyType分配给 []interface{}。

## 判断 interface 变量存储的是哪种类型
一个 interface 可被多种类型实现，有时候我们需要区分 interface 变量究竟存储哪种类型的值？类型断言提供对接口值的基础具体值的访问
```
t := i.(T)
```
该语句断言接口值i保存的具体类型为T，并将T的基础值分配给变量t。如果i保存的值不是类型 T ，将会触发 panic 错误。为了避免 panic 错误发生，可以通过如下操作来进行断言检查
```
t, ok := i.(T)
```
断言成功，ok 的值为 true,断言失败 t 值为T类型的零值,并且不会发生 panic 错误。
```
func main() {
	var i interface{}
	i = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	i = 100
	t, ok := i.(int)
	fmt.Println(t, ok)

	t2 := i.(string) //panic
	fmt.Println(t2)
}
```

## Type switch
还有一种方便的方法来判断 interface 变量的具体类型，那就是利用 switch 语句。如下所示：
```
func Print(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("type is string,value is:%v\n", i.(string))
	case float64:
		fmt.Printf("type is float32,value is:%v\n", i.(float64))
	case int:
		fmt.Printf("type is int,value is:%v\n", i.(int))
	}
}
func main() {
	var i interface{}
	i = "hello"
	Print(i)
	i = 100
	Print(i)
	i = 1.29
	Print(i)
}
```
灵活高效的 interface 动态类型,使 Go 语言在保持强静态类型的安全和高效的同时，也能灵活安全地在不同相容类型之间转换

## 参考
[Golang print nil](https://stackoverflow.com/questions/32318583/golang-print-nil)  
[InterfaceSlice](https://github.com/golang/go/wiki/InterfaceSlice)