
包管理工具
*  推荐使用[glide](https://github.com/Masterminds/glide)
* [Golang 依赖包管理工具对比](https://ieevee.com/tech/2017/07/10/go-import.html)

语法
* 如果一个名字以大写字母开头，那么它就是已导出的
* 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
* := 结构不能在函数外使用
* int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
* Go 在不同类型的项之间赋值时需要显式转换
* defer
    * defer 定义函数的参数为，为定义时候的值
        ```
        func a() {
            i := 0
            defer fmt.Println(i)
            i++
            return
        }
        // print 0

        func b(i int) {
	        defer func() {
		        fmt.Println(i)
	        }()
	        i++
        }
        b(1)
        // print 1
        ```
    * defer 定义多个函数时，是后进新出的执行顺序
        ```
        func b() {
            for i := 0; i < 4; i++ {
                defer fmt.Print(i)
            }
        }
        prints 3210
        ```
    * defer 可以对函数返回值进行操作
        ```
        func c() (i int) {
            defer func() { i++ }()
            return 1
        }
        // print 2
        ```

资源
* [golang fmt格式“占位符”](https://studygolang.com/articles/2644)