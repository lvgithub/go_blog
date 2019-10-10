
# 为什么在Go中有两种方式声明变量，有什么区别和使用哪些？

学习到声明变量语法的时候，有个疑问。Go作为一门语法非常简洁的语言，为什么一个简单的变量声明方式，会允许有两种不同方式同时存在。

* 声明变量
```
var name = "liuwei"
var sex int
```

* 声明变量(简写)
```
name := "liuwei"
sex := 1
```

既然有简写，为什么Go语言在设计同时存在两种不同的方式，对初学者来说肯定会有疑惑？

**两种变量之间区别**

声明变量 var 关键字是必须的，语意非常清晰、简短。
1. 批量声明变量
2. 允许不为变量赋值
    ```
    var  (
        name string
        sex int = 1
        isCoder = true
    )
    ```

简写的定义方式有什么优势呢，
1. 在 if、for 等语句很方便，简短的定义局部变量，非常符合Go语法清晰的设计原则
    ```
    for idx, value := range array {
        // Do something
    }

    if num := runtime.NumCPU(); num > 1 {
        // Do something
    }
    ```

2. 允许变量重新声明
    ```
    var name = "myfile.txt"

    fi, err := os.Stat(name)  first declared
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name, fi.Size(), "bytes")

    data, err := ioutil.ReadFile(name) 

    if err != nil {
        log.Fatal(err)
    }
    ```
    
允许重新声明变量，这大大避免了我们写代码的时候，冗余变量名称，对代码简洁起到了非常大的帮助。
