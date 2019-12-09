golang中你不知道的 string

字符串对于一篇博客文章来说似乎太简单了，但是一个简单的东西想用好，其实也不容易。

## 遍历字符串
```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const sample = "我爱golang"
	for i := 0; i < len(sample); i++ {
		runeValue, _ := utf8.DecodeRuneInString(sample[i:])
		fmt.Printf("position:%v, value:%c \n", i, runeValue)
	}
}
```
```
position:0, value:我 
position:1, value:� 
position:2, value:� 
position:3, value:爱 
position:4, value:� 
position:5, value:� 
position:6, value:g 
position:7, value:o 
position:8, value:l 
position:9, value:a 
position:10, value:n 
position:11, value:g
```
输出的是每一个字节。
如果想输出字符呢？
```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const sample = "我爱golang"
	for i, s := 0, 0; i < len(sample); i = i + s {
		runeValue, size := utf8.DecodeRuneInString(sample[i:])
		fmt.Printf("position:%v, value:%c \n", i, runeValue)
		s = size
	}
}
```
输出
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
还有更方便的吗？
```
package main

import (
	"fmt"
)

func main() {
	const sample = "我爱golang"
	for key, v := range sample {
		fmt.Printf("position:%v, value:%c \n", key, v)
	}
}
```
输出
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

* golang字符串是由字节构成的，因此索引它们会产生字节，而不是字符
* 通过unicode、utf-8等编码方式,从[]byte中解码字符串

## 字符串的长度
```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "我爱golang"
	fmt.Println(len(str), "len bytes")
	fmt.Println(utf8.RuneCountInString(str), "len characters") // 更快
	fmt.Println(len([]rune(str)), "len characters")
}
```
输出
```
12 len bytes
8 len characters
8 len characters
```
* golang对字符串的底层处理，采用的是[]byte，实际存储的值是一个uint8类型的;
* 采用UTF-8编码英文字符，每个字符只占用一个byte,而中文需要占用3个byte，因此长度是12;
* utf8会自动判断每个字符编码占用了几个byte,很清晰的展示了 Golang 对 string 的处理原理;
* for range string 每次迭代会解码一个 utf-8 编码的字符。

## 修改字符串
```
package main

import "fmt"

func main() {
	str := "golang"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
}
```
输出
```
colang
```



## 拼接字符串

```
package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "我爱"
	str2 := "golang"
	// 该方案每次合并会创建一个新的字符串
	fmt.Println(str + str2)

	// 该方案更更快，直接连接底层的 []byte
	var buffer bytes.Buffer
	buffer.WriteString(str)
	buffer.WriteString(str2)
	fmt.Println(buffer.String())
}

```
WriteString更快的原因,见源码直接底层[]byte连接
```
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(len(s))
	if !ok {
		m = b.grow(len(s))
	}
	return copy(b.buf[m:], s), nil
}
```
## 参考
* https://blog.golang.org/strings