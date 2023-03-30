package main

import (
	"fmt"
	"time"
)

/*
我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。
Go 内置的 定时器 和 打点器 特性让这些变得很简单。 我们会先学习定时器，然后再学习打点器。*/

func main() {
	/*定时器表示在未来某一时刻的独立事件。 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。 这里的定时器将等待 2 秒。*/
	timer1 := time.NewTimer(2 * time.Second)
	//<-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
	//大白话就是 timer1.C 想要往外发送值 但是本身内却没有值 于是一直堵塞 知道有人给他发送了值 这个人就是定时器 fire 小火枪
	<-timer1.C
	fmt.Println("timer 1 fired")
	/*如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。
	使用定时器的原因之一就是，你可以在定时器触发之前将其取消。 例如这样。*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	//timer2.C 会一直堵塞 因为本身空 且每人给他发送值 他会一直堵塞到有人给他发送值
	//堵塞等待过程中 意外发生了 调用了timer2.Stop() 也就是timer2停止了
	//最后一句是反验证 我多等等你timer2（给了4s也不行 再多给也不行） 希望你能触发 实际是还是没触发
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stoped")
	}
	//给 timer2 足够的时间来触发它，以证明它实际上已经停止了。
	time.Sleep(4 * time.Second)
}
