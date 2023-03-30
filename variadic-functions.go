package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4, 5)
	nums := []int{11, 22, 33, 44, 55}
	//sum(nums) //错误写法
	sum(nums...)
	//如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)。
}

//这个函数接受任意数量的 int 作为参数。
//变的是参数数量
func sum(nums ...int) {
	fmt.Print(" ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
