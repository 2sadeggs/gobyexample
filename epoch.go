package main

import (
	"fmt"
	"time"
)

/*一般程序会有获取 Unix 时间 的秒数，毫秒数，或者微秒数的需求。来看看如何用 Go 来实现。*/

func main() {

	/*分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。*/
	now := time.Now()
	//secs := time.Now().Second() //思路错了 是直接用now的秒  而不是再获得一个秒
	secs := now.Second()
	nanos := now.Nanosecond()
	fmt.Println(now)

	/*注意 UnixMillis 是不存在的，所以要得到毫秒数的话， 你需要手动的从纳秒转化一下。*/
	/* 1秒(s) ＝1000毫秒(ms)
	1毫秒(ms)＝1000微秒 (us)
	1微秒(us)＝1000纳秒 (ns)
	1纳秒(ns)＝1000皮秒 (ps)*/
	//明显毫秒跟纳秒之间差两个级别 也就是6个0  很明显 1ms=1000000ns  于是纳秒/1000000
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	/*你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。*/
	/*	fmt.Println(time.Unix(secs, 0))
		fmt.Println(time.Unix(0, nanos))*/
	//错误写法 导致下面错误
	//错误原因是因为中文版的教程就错了  实际上英文版教程没错 两者不同  所以有能力一定阅读英文原文
	//.\epoch.go:31:24: cannot use secs (variable of type int) as type int64 in argument to time.Unix
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

/*2023-04-01 20:58:49.5104818 +0800 CST m=+0.006769601
49
510
510481800
2023-04-01 20:58:49 +0800 CST
2023-04-01 20:58:49.5104818 +0800 CST
*/
