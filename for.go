package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//最基础的方式，单个循环条件。

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//经典的初始/条件/后续 for 循环。

	for {
		fmt.Println("loop")
		fmt.Println("loop.")
		fmt.Println("loop..")
		break
	}
	//不带条件的 for 循环将一直重复执行， 直到在循环体内使用了 break 或者 return 跳出循环。

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println("hahaha..bingo")
			continue //如果没有continue 即便是命中也会执行下边的打印，如果加上continue则命中的值不再往下执行，直接下一个循环
			//简言之 continue跳出当前循环进入下一个循环
		}
		fmt.Println(n)
	}
	//你也可以使用 continue 直接进入下一次循环。
}
