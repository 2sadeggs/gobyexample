package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println(a)
	//go会自动推算已经有初始值的变量类型
	var aa string = "hahaha" //这种写法比较规范 初学者推荐
	fmt.Println(aa)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	//声明后未给初始化值 会默认初始化为对应变量类型的零值

	f := "short" // := 语法为声明并初始化变量的简写 for循环或者if case 逻辑中常见
	fmt.Println(f)
}
