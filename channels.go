package main

import "fmt"

/*
channels是连接多个协程的管道
你可以从一个协程中将值发送到管道 然后在另一个协程中从管道接受
*/

func main() {
	//使用 make(chan val-type) 创建一个新的通道。 通道类型就是他们需要传递值的类型。
	messages := make(chan string)
	//使用 channel <- 语法 发送 一个新的值到通道中。 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
	go func() {
		messages <- "ping"
	}()
	//使用 <-channel 语法从通道中 接收 一个值。 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
	//msg := <-messages
	//fmt.Println(msg)

	//总结 从子协程传递值到管道 无缓冲管道 会堵塞 知道在主协程中取出值 堵塞完成 然后打印出来

	//再取一遍 看发生什么
	//msg2 := <-messages
	//fmt.Println(msg2)
	//fatal error: all goroutines are asleep - deadlock! 死锁了

	//取出值但是我不要了 看看能不能正常执行下一步
	_ = <-messages
	//	_ := <-messages  这个是想当然的错误写法 因为 _ 不是变量 所以也就不需要 := 运算了
	fmt.Println("从管道取出了值 但是直接丢弃了")
}
