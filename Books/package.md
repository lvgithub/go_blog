编写第一个包，入门Go的包管理方式

## 介绍
* Go 语言的源码复用建立在包（package）基础之上
* main() 函数所在的包（package）叫 main
* main 包想要引用别的代码，必须同样以包的方式进行引用

## 特定
* 一个目录下的同级文件属于一个包
* 包名可以与目录名不同（不建议）
* 包名为main的包为应用程序入口，一个程序必须有一个main包，只能有一个。
## 构建
* 创建包目录
```  
mkdir -p lvgithub.com/greet  
vim lvgithub.com/greet/greet.go
```

```
// greet.go
package greet

import "fmt"

// Say hello
func Say() {
	fmt.Printf("Hello Golang")
}
```
* 创建主文件  
```
vim /hello.go
```
```
// hello.go
package main

import "lvgithub.com/greet"

func main() {
	greet.Say()
}
```
* 运行  
```
 go run hello.go
 // Hello Golang
```

[源码地址](../Code/firstLib/hello.go)

## 参考
[Go语言包（package）](http://c.biancheng.net/golang/package/)