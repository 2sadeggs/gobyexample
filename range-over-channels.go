package main

import "fmt"

/*遍历通道*/

func main() {
	//	定义一个有两个缓存的channel 类型string
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//range 迭代从 queue 中得到每个值。 因为我们在前面 close 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。
	for elem := range queue {
		fmt.Println(elem)
	}
	//这个例子也让我们看到，一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。

	//这时候我再取一个值看看是什么
	oneMore := <-queue
	fmt.Println("oneMore:", oneMore) //结果是空 可能是空字符串

	//下边改变思路 改为int型channel
	q := make(chan int, 2)
	//qq := make(chan int, 3)
	q <- 111
	q <- 222
	close(q)
	/*	for j := 0; j < 3; j++ {
		fmt.Println("qq: ", <-q)
	}*/
	for elem := range q {
		fmt.Println(elem)
	}
	aaa := <-q
	fmt.Println("oneMoreAgain:", aaa)
	//这次的结果是0 哈哈 符合预期 管道关闭后取出所有值后还能继续取值 且取出的都是管道对应类型的零值
}
