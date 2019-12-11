
通过 Once学习 Go 的内存模型

Once 官方描述 Once is an object that will perform exactly one action,即 Once是一个对象,它提供了保证某个动作只被执行一次功能，最典型的场景就是单例模式。

## 单例模式
```
package main

import (
	"fmt"
	"sync"
)

type Instance struct {
	name string
}

func (i Instance) print() {
	fmt.Println(i.name)
}

var instance Instance

func makeInstance() {
	instance = Instance{"go"}
}

func main() {
	var once sync.Once
	once.Do(makeInstance)
	instance.print()
}
```
once.Do 中的函数只会执行一次，并且可以保证 once.Do 返回时，传入Do的函数已经执行完成。（多个 goroutine 同时执行 once.Do 的时候，可以保证抢占到 once.Do 执行权的 goroutine 执行完 once.Do 后，其他 goroutine 才得到返回 ）

## 源码
源码很简单，但是这么简单不到20行的代码确能学习到很多知识点，非常的强悍。
```
package sync

import (
	"sync/atomic"
)

type Once struct {
	done uint32
	m    Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```
这里的几个重点知识：
1. Do 方法为什么不直接 o.done == 0 而要使用 atomic.LoadUint32(&o.done) == 0  
2. 为什么 doSlow 方法中直接使用 o.done == 0
3. 既然已经使用的Lock, 为什么不直接 o.done = 1， 还需要 atomic.StoreUint32(&o.done, 1)

先回答第一个问题？如果直接 o.done == 0，会导致无法及时观察 doSlow 对 o.done 的值设置。具体原因可以参考 [Go 的内存模型](https://golang.org/ref/mem) ,文章中提到：
```
Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.

To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages.
```
大意是 当一个变量被多个 gorouting 访问的时候，必须要保证他们是有序的（同步），可以使用 sync 或者 sync/atomic 包来实现。用了 LoadUint32 可以保证 doSlow 设置 o.done 后可以及时的被取到。

再看第二个问题，可以直接使用 o.done == 0 是因为使用了 Mutex 进行了锁操作，o.done == 0 处于锁操作的临界区中，所以可以直接进行比较。

相信到这里，你就会问到第三个问题 atomic.StoreUint32(&o.done, 1) 也处于临界区，为什么不知直接通过 o.done = 1 进行赋值呢？这其实还是和内存模式有关。**Mutex 只能保证临界区内的操作是可观测的** 即只有处于o.m.Lock() 和 defer o.m.Unlock()之间的代码对 o.done 的值是可观测的。那这是 Do 中对 o.done 访问就可以会出现观测不到的情况，因此需要使用 StoreUint32。

到这里是不是发现了收货了好多，还有更厉害的。
我们再看看为什么 dong 不使用 uint8或者bool 而要使用 uint32呢？
```
type Once struct {
    // done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/x86),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	m    Mutex
}
```
目前能看到原因是，atomic 包中没有提供 LoadUint8 、LoadBool 的操作。  
然后看注释，我们发现更为深奥的秘密：我们看他的注释：它提到一个重要的概念 **hot path**，即 Do 方法的调用会是高频的，而每次调用访问 done，done位于结构体的第一个字段，可以通过结构体指针直接进行访问（访问其他的字段需要通过偏移量计算就慢了）


