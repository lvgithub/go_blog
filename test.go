package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func main() {
	println("hello")
	println("world")
	fmt.Print("go\n")
}

func read(n int) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)

	v := count

	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()
}

func write(n int) {
	rw.Lock()
	defer rw.Unlock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)

	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
}
