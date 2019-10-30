Go中什么时候函数参数按值传递？

C语言系列的所有语言，函数的参数总是按值传递也就是说函数总能得到参数的副本，就像有一个赋值语句将值分配给参数意义。

其实指针传递也是赋值一个指针的副本，不是他指向的数据。所以你能在函数里面通过指针改变数据的值就是这个原因。因为指针是一个指向数据的地址，副本还是这个地址。

有两个比较特殊的数据结构就是 Map 和 Slice。它们的行为和指针很像，它们包含一个指向基础数据的一个指针，可以查看[深入理解 Slice](https://github.com/lvgithub/go_blog/blob/master/Books/slice.md)。
所以 Map 和 Slice 作为参数的时候，不会复制他们所指向的数据，而且在函数中可以直接改变数据。

如果函数的参数是一个 interface ,interface 值包含 struct 将会复制整个 struct，如果包含的是指针，那也不会拷贝指针指向的数据。



参考：[When are function parameters passed by value?](https://golang.org/doc/faq#pass_by_value)



