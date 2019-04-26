package main

import (
	"log"
	"time"
	"sync"
)

func doSomething(jobId int, group *sync.WaitGroup)  {
	log.Printf("start jobid:[%v]\n", jobId)
	time.Sleep(3 * time.Second)
	log.Printf("end jobid:[%v]\n", jobId)
	group.Done()
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(3)

	for i:=0;i<3;i++  {
		//使用 goroutine 运行
		go doSomething(i, &waitgroup)
	}
	 waitgroup.Wait()
}