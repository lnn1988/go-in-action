# 并发模式

## 主要内容
* 控制程序的生命周期
* 管理可以复用的资源池
* 创建可以处理任务的 goroutine 池

## runner
runner 包用于展示如何使用通道来监视程序的执行时间。如果执行时间过长，也可以使用 runner 包进行终止程序运行。

