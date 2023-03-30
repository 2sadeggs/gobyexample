package main

import (
	"fmt"
	"time"
)

/*
我们可以使用通道来同步协程之间的执行状态。
这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。
如果需要等待多个协程，WaitGroup 是一个更好的选择。
*/
//定义一个worker函数 就是干活的 模拟干活过程 入参是一个channel bool类型的channel

func worker(done chan bool) {
	fmt.Println("hard working...")
	time.Sleep(2 * time.Second) // 模拟工作 也可以说划水
	fmt.Println("done")

	done <- true //发送一个值来通知我们已经完工啦。
	//虽然是一个入参 但是感觉不是入的方向 二是出的方向 告诉外边我们完成工作啦

}
func main() {
	//定义一个channel 不再用关键字done 以免混淆
	over := make(chan bool, 1)
	//运行一个 worker 协程，并给予用于通知的通道。
	go worker(over)

	//程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	//可以将划水时间提升至2秒 提升演示效果
	<-over

	/*
		如果你把 <- done 这行代码从程序中移除， 程序甚至可能在 worker 开始运行前就结束了。
	*/
}
