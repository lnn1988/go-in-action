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