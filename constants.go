package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n //常数表达式可以执行任意精度的运算
	fmt.Println(d)
	//3e20  3*10的20次方 5e3 就是5乘以10的三次方 也就是5000 验证
	fmt.Println(5e3)

	fmt.Println(int64(d)) //数值型常量没有固定的类型 直到被给定某个类型 比如显示类型转换

	fmt.Println(math.Sin(n))
	//常量数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型 比如Math.Sin需要一个float64的参数，n会自动确定为float64
	//构造一个函数试一下 注意：不是隐式转 二是显示转 此例就是用string显示转为string类型
	Testconstant(string(n))

}
func Testconstant(s string) {
	fmt.Println("constant now is type of string")
}

/*
 go支持 字符 字符串 布尔和数值类型常量
 const语句可以出现在任何var语句可以出现的地方 也就是两者等权
 常数表达式可以执行任意精度的运算
 数值型常量没有固定的类型 直到被给定某个类型 比如显示类型转换
 常量数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型 比如Math.Sin需要一个float64的参数，n会自动确定为float64
*/
