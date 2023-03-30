package main

import (
	"fmt"
	"time"
)

//定义一个函数 循环输出
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}
func main() {
	f("direct") // main goroutine 直接调用

	go f("goroutine") //子 goroutine

	go func(msg string) { // 匿名函数 goroutine
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) //用于等待子goroutine执行完
	//time.Sleep(3 * time.Second)
	//更好的方法是使用 WaitGroup
	fmt.Println("done")
	/*
	      当我们运行这个程序时，首先会看到阻塞式调用的输出，然后是两个协程的交替输出。
	   这种交替的情况表示 Go runtime 是以并发的方式运行协程的。
	*/
	//执行结果顺序一直是 direct going goroutine 明显不是顺序 但是并发随机性体现也不是很明显
}
