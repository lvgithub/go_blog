为什么要使用 goroutines 取代 threads

## 介绍
goroutines 和 threads 都是为了并发而生。准备的说，并不能说 goroutines 取代 threads。因为其实 goroutines 是建立在一组 threads 之上。将多路并发执行的能力复用在 threads 组上。当某一个 goroutine 需要运行时，会自动将它挂载到一个 thread 上。而且这一系列的调度对开发者是黑盒，无感知的。


## goroutines 对比 threads 的优势
* 开销非常便宜，每个 goroutine 的堆栈只需要几kb,并且可以根据应用的需要进行增长和缩小。
* 一个 thread 可以对应多个 goroutine，可能一个线程就可以处理上千个 goroutine。如果该线程中有 goroutine 需要被挂起，等待用户输入，那么会将其他需要运行的 goroutine 转移到一个新的 thread 中。所有的这些都被抽象出来，开发者只要面对简单的 API 就可以使用。

## 使用 goroutine
```
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine\r")
}
func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function\r")
}

```
细心的同学一定会问到为什么需要执行 	time.Sleep(1 * time.Second) 这句话。因为：
* 当创业一个 goroutine 的时候，会立即返回，执行下语句，而且忽悠所有 goroutine 的返回值；
* 如果主 goroutine 退出，则其他任何 goroutine 将不会被执行；
* 如果你注释 Sleep 这句话，再运行一次，将会看不到 hello goroutine输出。

## 多个 goroutine
```
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'c'; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
```
输出为：
```
a 1 b 2 c 3 d 4 e 5 main terminated
```
下面表格描述多个 goroutine 执行的时序,可以看出多个 goroutine 是同时进行的。
numbers goroutine
|0ms|500ms|1000ms|1500ms|2000ms|2500ms|
|---|-----|-----|-----|-----|-----|
|   | 1   |2    | 3   | 4   |   5 |

alphabets goroutine
|0ms|400ms|800ms|1200ms|1600ms|2000ms|
|---|-----|-----|-----|-----|-----|
|   | a   |b    | c   | d   |   e |

main goroutine
|0ms|400ms|500ms|800ms|1000ms|1200ms|1500ms|1600ms|2000ms|2500ms|3000ms|
|---|-----|-----|-----|------|------|------|------|------|------|------|
|   | a   |1    | b   | 2    |  c   | 3    |d     |4(e)  |  5   |      |



