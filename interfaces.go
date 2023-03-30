package main

import (
	"fmt"
	"math"
)

//方法签名的集合叫做：_接口(Interfaces)_。

//定义一个接口 几核
type geometry interface {
	//该接口有两个方法 面积和周长
	area() float64
	perim() float64
}

//定义两个结构体矩形和圆形 （rect和circle）rect和circle 来实现这个接口 得实现接口的所有方法
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//注意 实现接口 实现全部接口方法 且！！记得原样返回接口的返回值 此例返回值是float64

//定义一个通用测量函数 入参为接口类型geometry
//也就是说本义是想测量所有几何图形的面积和周长
//
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	//初始化一个矩形和圆形
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	//r和c 是rect和circle类型 且该两个类型都实现了 geometry接口
	//如果变量是一个接口类型 那么此变量就可以直接调用接口中的方法
	//如例子中的g.area()和g.perim()
	measure(r)
	measure(c)
}
