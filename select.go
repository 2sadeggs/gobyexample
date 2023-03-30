package main

import (
	"fmt"
	"time"
)

/*
Go 的 选择器（select） 让你可以同时等待多个通道操作。 将协程、通道和选择器结合，是 Go 的一个强大特性。
*/

func main() {
	//初始化两个无缓冲channel
	//两个string类型的channel
	c1 := make(chan string)
	c2 := make(chan string)
	/*
	   各个通道将在一定时间后接收一个值， 通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）。
	*/
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	//我们使用 select 关键字来同时等待这两个值， 并打印各自接收到的值。
	//注意是同时等待两个值 也就是两个case都会执行 但是执行结果顺序不固定 有可能 one two 也有可能two one
	//如果想顺序执行可以调整c1 和 c2 的sleep时间 让c1少睡会 1s c2多睡会 2s  这样结果总是 one two
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("received: ", msg2)
		}

	}
}
