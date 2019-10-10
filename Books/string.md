深入理解 string

字符串对于一篇博客文章来说似乎太简单了，但是要想很好地使用它们，不仅需要了解它们是如何工作的，而且还需要了解他的工作原理。
会关注到这点，其实是在开发过程中的使用for range 遍历字符串的时候，发现取到的结果便不是字符，这于其他语言，如C语言有很大的实现区别，因此就查阅了官方文档，进行深入的了解。

开门见山，先说结论吧。
* golang字符串是由字节构成的，因此索引它们会产生字节，而不是字符
* 通过unicode、utf-8等编码方式,从[]byte中解码字符串

看看字符串的长度
```
const text = "helloboy,你好"
fmt.Println(len(text))// 15
```
上面的代码结果为什么是15而不是11
* golang对字符串的底层处理，采用的是[]byte，实际存储的值是一个uint8类型的
* 采用UTF-8编码英文字符，每个字符只占用一个byte,而中文需要占用3个byte，因此长度是15

for range string 是什么样的结果呢？
```
const sample = "我爱golang"
for key, v := range sample {
    fmt.Printf("position:%v, value:%c \n", key, v)
}
```
输出结果:
```
position:0, value:我 
position:3, value:爱 
position:6, value:g 
position:7, value:o 
position:8, value:l 
position:9, value:a 
position:10, value:n 
position:11, value:g 
```
结果特点很明显，对于中文的 position 每循环一次递增步长为3，英文字符为1。很困惑把，看下面的代码你就能解惑了。
```
const sample = "我爱golang"
for i, w := 0, 0; i < len(sample); i += w {
    runeValue, width := utf8.DecodeRuneInString(sample[i:])
    fmt.Printf("position:%v, value:%c \n", i, runeValue)
    w = width
}
```
结果为:
```
position:0, value:我 
position:3, value:爱 
position:6, value:g 
position:7, value:o 
position:8, value:l 
position:9, value:a 
position:10, value:n 
position:11, value:g
```
两个结果一样，utf8会自动判断每个字符编码占用了几个byte,很清晰的展示了 Golang 对 string 的处理原理。而 for range string 每次迭代会解码一个 utf-8 编码的字符。
这个程序使用 utf8 类库进行处理，相当于上面的 for range string 示例。

## 参考
* https://blog.golang.org/strings