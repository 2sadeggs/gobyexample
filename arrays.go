package main

import "fmt"

func main() {
	//定义一个数组a 存放5个int
	var a [5]int
	fmt.Println("empty: ", a)
	//数组默认值是零值，对于 int 数组来说，元素的零值是 0。

	//我们可以使用 array[index] = value 语法来设置数组指定位置的值， 或者用 array[index] 得到值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	//内置函数 len 可以返回数组的长度。
	fmt.Println("len: ", len(a))

	//使用这个语法在一行内声明并初始化一个数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	//数组类型是一维的，但是你可以组合构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i==j", i, j)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	//注意，使用 fmt.Println 打印数组时，会按照 [v1 v2 v3 ...] 的格式打印。
}
