# 使用 go modules 进行包管理

## 背景
2018年3月8日，在Go社区开启了在Go工具链中支持包版本的[讨论](https://github.com/golang/go/issues/24301)，Go 1.11对此处提出的版本化[草案](https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md)进行了初步支持。

关于Go Modules的发布，社区还引起了很大的争端，有兴趣可以查看[《关于Go Module的争吵》](https://zhuanlan.zhihu.com/p/41627929)

## 介绍
* 从 Go 1.11 开始，Go 允许在 $GOPATH/src 外的任何目录下使用 go.mod 创建项目
* 为了兼容性，Go 命令仍然在旧的 GOPATH 模式下运行，需要使用go modules ，执行命名 **export GO111MODULE=on**开启

## GO111MODULE 模式
* **auto** 模式，项目在$GOPATH/src内使用$GOPATH/src的依赖包，在$GOPATH/src外，使用 go.mod 里require的包
* **on** 模式，1.12后，无论在$GOPATH/src里还是在外面，都会使用go.mod 里require的包
* **off** 模式，就是老规矩
## 使用
* 创建工程
```
$ mkdir -p $GOPATH/blog/Code/mod
$ cd $GOPATH/blog/Code/mod
```
* 初始化一个modules
```
$ go mod init lvgithub.com/mod
go: creating new go.mod: module lvgithub.com/mod
```
* 编写代码
```
$ vim hello.go

package main
import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

* 编译
```
## go 会自动查找代码中的包，下载依赖包，并且把具体的依赖关系和版本写入到go.mod和go.sum文件中
$ go build ./hello.go

go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
```
* 执行
```
$ ./hello.go 
If a program is too slow, it must have a loop.
```
## VSCode支持
如果你的VSCode编辑器对使用go modules的包无法找到，请参考[Go modules support in Visual Studio Code](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code)进行配置

## 依赖包下载到哪里
* 成功运行本示例后，可以在$GOPATH/pkg/mod下发现**rsc.io/quote@v1.5.1**包
* 如果你修改版本后，会生成**rsc.io/quote@[version]** 包，也就是说一个包在pkg/mod下可以拥有不同的版本    

## 依赖地址改动
```
$ vim go.mod
# 用 replace 替换包
replace golang.org/x/text => github.com/golang/text latest
```
## 参考
[Proposal: Versioned Go Modules](https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md)
[Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)

[测试代码地址](../Code/mod/hello.go)