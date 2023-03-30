package main

import "fmt"

/*
关闭 一个通道意味着不能再向这个通道发送值了。 该特性可以向通道的接收方传达工作已经完成的信息。
*/
//关闭通道时 通道里边可能还有值 并且还能取出来 接下来验证

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//在这个例子中，我们将使用一个 jobs 通道，
	//将工作内容， 从 main() 协程传递到一个工作协程中。
	//当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。
	go func() { //第一次少了for循环 所以报错死锁
		for {
			j, more := <-jobs //more 更常用的变量名为ok  表示结束值是否成功
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true //表明我干完了
				return       //第一次忘记return了 结果报错
			}
		}
	}()
	//使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job ", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	//done <- true //使用前面学到的通道同步方法等待任务结束。//第一次的错误写法 不是再接受值了 二是把值取出
	<-done
	//如果不取出 朱谢成退出 子协程可能接受值不全 也就是可能接受到一半就退出了
}
