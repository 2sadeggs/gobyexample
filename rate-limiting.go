package main

import (
	"fmt"
	"time"
)

/*速率限制 是控制服务资源利用和质量的重要机制。 基于协程、通道和打点器，Go 优雅的支持速率限制。*/

func main() {
	/*首先，我们将看一个基本的速率限制。 假设我们想限制对收到请求的处理，我们可以通过一个渠道处理这些请求。*/
	//模拟创建5个请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	//然后关闭管道
	close(requests)
	//limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	limiter := time.Tick(200 * time.Millisecond)
	//通过在每次请求前阻塞 limiter 通道的一个接收， 可以将频率限制为，每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	/*
			有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
		我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	*/
	burstyLimiter := make(chan time.Time, 3) //乍一看好像是通过缓冲实现的爆发 注意channel的类型为time.Time
	//填充通道，表示允许的爆发（bursts）。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	//每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	//所以三个并发最理想的情况是三个并发一下子过来占领了全部burstyLimiter 正好缓冲3个
	//关键是一下子 也就是并发过来
	/*现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成。*/
	/*现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成。*/
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	//关闭channel
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	/*运行程序，我们看到第一批请求意料之中的大约每 200ms 处理一次。*/
	/*第二批请求，由于爆发（burstable）速率控制，我们直接连续处理了 3 个请求，
	然后以大约每 200ms 一次的速度，处理了剩余的 2 个请求。*/

	/*
		PS D:\GolandProjects\gobyexample> go run .\rate-limiting.go
		request 1 2023-03-30 17:05:07.5620772 +0800 CST m=+0.207425501
		request 2 2023-03-30 17:05:07.7663488 +0800 CST m=+0.411696601
		request 3 2023-03-30 17:05:07.9702351 +0800 CST m=+0.615582401
		request 4 2023-03-30 17:05:08.1576175 +0800 CST m=+0.802964301
		request 5 2023-03-30 17:05:08.3594085 +0800 CST m=+1.004754801
		request 1 2023-03-30 17:05:08.3597001 +0800 CST m=+1.005046401
		request 2 2023-03-30 17:05:08.3597001 +0800 CST m=+1.005046401
		request 3 2023-03-30 17:05:08.3597001 +0800 CST m=+1.005046401
		request 4 2023-03-30 17:05:08.5621907 +0800 CST m=+1.207536501
		request 5 2023-03-30 17:05:08.7652311 +0800 CST m=+1.410576401
	*/
}
