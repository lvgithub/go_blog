
构建并运行Go程序

## 前置条件
## 重要的环境变量 GOPATH
**GOPATH**，项目的工作目录，是Go编译器在构建Go应用程序时用来搜索依赖项的：
* **$GOPATH/src** ，代码保存的目录
* **$GOPATH/bin** ，工程经过 go build、go install 、go get 等指令后，产生的二进制可执行文件放目录
* **$GOPATH/pkg** ，生成的中间缓存文件(.a)保存的目录


**GOROOT**，Go 的安装路径  
**GOBIN**，go install编译存放路径，为空时可执行文件放在各自GOPATH目录的bin文件夹中

**在最新的主要版本Go v1.11中，GOPATH不再是强制性的**，在这个新版本中引入了[go modules](https://github.com/golang/go/wiki/Modules)
## 构建我们的第一个Go程序了
创建目录
```
mkdir -p $GOPATH/src/hello
```
创建hello.go
```
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```
编译
* 编译完成的可执行文件会保存在 $GOPATH/bin 目录下。
```
go install hello.go // 将可执行文件或库文件安装到 $GOPATH/bin
go build hello.go // 在生成可执行文件在当前目录
```
运行
```
$GOPATH/bin/hello
```