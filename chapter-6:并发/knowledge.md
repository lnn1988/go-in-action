# 竞争状态

如果多个 goroutine 在没有互相同步的情况下，访问某个共享的资源，试图同时读或者写这个资源，就处于互相竞争的状态。
对一个共享资源的读或者写操作必须是原子化的。
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
)

func incCounter(id int)  {
	defer wg.Done()

	for count:=0;count<2;count++ {
		value := counter
		//让出此线程，并把当前 goroutine 放到队列
		runtime.Gosched()
		value ++
		counter = value
	}
}

func main() {
	wg.Add(2)
	
	go incCounter(1)
	go incCounter(2)
	
	wg.Wait()
	fmt.Println("final counter: ", counter)
}
```

一种修正代码、消除竞争状态的办法是，使用 Go 语言提供的锁机制，锁住共享资源，保证 goroutine 的同步状态。

## 锁住共享资源
Go 语言提供了传统的同步 goroutine 的机制，就是对共享资源加锁。
如果要顺序访问一个整形变量或者一段代码，**atomic**和**sync**标准库的函数提供了很好的解决方案。
```
//原子操作
atomic:
    atomic.AddInt64()
    atomic.LoadInt64()
    atomic.StoreInt64()

// 互斥锁
sync:
    sync.mutex()
```
## 原子操作
atomic 包提供了 AddInt64 函数，这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作。
另外两个有用的函数是 LoadInt64 和 StoreInt64。这两个函数提供了一种安全的读写一个整型值的方式。

## 互斥锁
另一种同步访问共享资源的方式是使用互斥锁(mutex)。
互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码区域。
```
package main

import (
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
	// 用来定义一段临界代码区
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
}

func incCounter(id int)  {
	defer wg.Done()

	for count := 0; count<2; count++ {
		// 用来标注，同一时只能允许一个 goroutine 进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value ++
			counter = value
		}
		// 释放互斥锁
		mutex.Unlock()
	}
}
```




























