package main

import "fmt"

/*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收，甚至是非阻塞的多路 select。
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	/*这是一个非阻塞接收的例子。
	如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。 否则，就直接到 default 分支中。*/
	select {
	case msg := <-messages:
		//因为此时messages里 没有值 因此 msg取不到  发生堵塞 故配合select机制 走了另外一个不阻塞分支
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	/*
	   一个非阻塞发送的例子，代码结构和上面接收的类似。
	   msg 不能被发送到 message 通道，因为这是 个无缓冲区通道，
	   并且也没有接收者，因此， default 会执行。
	*/
	msg := "hi"
	select {
	case messages <- msg:
		//显示是msg把值发送给messages了
		//但是因为messages是无缓冲的chan 此时没人从messages中取值 因此发生堵塞 配合select机制 走其他分支
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	/*
		我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
		这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signals", sig)
	default:
		fmt.Println("no activity")
	}
	//结果很有趣 三个分支都执行了 比上上个例子的两个分支还多
}
