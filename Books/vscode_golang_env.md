
VS Code Golang工具链插件 vscode-go

[插件仓库](https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on)

## 插件
* [gocoder 代码自动补全提示工具](https://github.com/nsf/gocode)
* [gopkgs 自动导入包工具](https://github.com/uudashr/gopkgs)

## 安装
1. 使用快捷键：**command+shift+P** 打开命令模式
2. 键入: **go:install/update tools**
3. 将所有插件都勾选上
4. 点击 OK 即开始安装

真的这么顺利吗？不是的，哈哈，安装过程中需要访问golang.org获取代码包，你懂的，这时候需要梯子。

## 梯子
如果你已经有梯子，直接在VS Code 设置代理,重新安装即可
```
"http.proxy": "http://proxy.agent.com"
```

## 官方梯子
如果没有呢？使用 [goproxy.io](https://goproxy.io/) 官方梯子进行安装
```
# 启用 go modules 功能
export GO111MODULE=on

# 设置 GOPROXY 环境变量
export GOPROXY=https://goproxy.io

# 安装 gocode
go get -u -v github.com/mdempsky/gocode
```
