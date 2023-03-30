package main

/*函数 是 Go 的核心。我们将通过一些不同的例子来进行学习它。*/

import "fmt"

func plus(a int, b int) int {

	return a + b
}

//当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	//通过 函数名(参数列表) 来调用函数
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
