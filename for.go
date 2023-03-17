package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		fmt.Println("loop.")
		fmt.Println("loop..")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println("hahaha..bingo")
			continue //如果没有continue 即便是命中也会执行下边的打印，如果加上continue则命中的值不再往下执行，直接下一个循环
			//简言之 continue跳出当前循环进入下一个循环
		}
		fmt.Println(n)
	}
}
