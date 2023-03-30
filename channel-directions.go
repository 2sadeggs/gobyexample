package main

import "fmt"

/*
语法
创建只读channel
var chanName <-chan chanType
说明
声明只读的 channel，只需要在 chan 关键字前面加上 <- 符号即可，此时 chanName 只能用于接收数据，而不能用于写数据。
所以又叫只读  chanName只能用于接收数据 而不能往外写数据 所以称只读
因为是从右边赋值 所以 <- 在左边好理解 在左边便表示 chan只能发送数据 所以 chanName只能用于接收
<- 在右边 表示 chan只能接收数据 所以对应的chanName只能发送 也就是网chan里发送
chanName只能往外写数据 所以称只写
创建只写channel
语法
var chanName chan<- chanType
说明
声明只写的 channel，只需要在 chan 关键字后面加上 <- 符号即可，此时 chanName 只能用于发送数据，而不能用于接收数据。

*/

/*
总而言之 言而总之 与直观感觉想法
我直观感觉 chan<-  明显表示陈能接收
<-chan 只能发送
实际情况想法 要加一层调用者或者初始化实例 这样才好理解 好几亿
*/
//This ping function only accepts a channel for sending values.
//It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
	//问 只能发送的chan pings 为什么还是能接收值msg
	//尝试从这个通道接收数据会是一个编译时错误。
}

//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	/*以上感觉双向channel自动转为单向channel了*/
}

/*
2023.03.29
推导错误记忆重新记忆

func ping(pings chan<- string, msg string)
还是还原直观感受 明显chan<- 表示只能接收数据的channel
后边把msg的值给pings 正常没有报错 pings <- msg

func pong(pings <-chan string, pongs chan<- string)
明显<-chan是一个只能发送的channel 后边的句子把pings 的值发给msg也没报错
msg := <-pings
msg 变量在函数中第一次使用 之前未声明 因此采用了短声明方式 :=
放弃只读只写的理解  有点绕
*/
