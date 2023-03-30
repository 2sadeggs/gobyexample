package main

import "fmt"

/*
Go支持对于结构体(struct)和接口(interfaces)的 嵌入(embedding) 以表达一种更加无缝的 组合(composition) 类型
*/

//定义一个base结构体
type base struct {
	num int
}

//为结构体绑定一个describe()方法 返回一个string
func (b base) describe() string {
	//zhuyi Sprintf()函数的返回值为string
	return fmt.Sprintf("base with num=%v", b.num)
}

//一个 container 嵌入 了一个 base. 一个嵌入看起来像一个没有名字的字段
type container struct {
	base
	str string
}

func main() {
	//当创建含有嵌入的结构体，必须对嵌入进行显式的初始化； 在这里使用嵌入的类型当作字段的名字
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	//我们可以直接在 co 上访问 base 中定义的字段, 例如： co.num.
	//直观感觉应该是 co.base.num 这么访问 感觉是语法糖
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	//或者，我们可以使用嵌入的类型名拼写出完整的路径。
	fmt.Println("also num:", co.base.num)

	/*
		由于 container 嵌入了 base，因此base的方法 也成为了 container 的方法。
		在这里我们直接在 co 上 调用了一个从 base 嵌入的方法。
	*/
	fmt.Println("describe:", co.describe()) //直观感觉 应该这么调用 co.base.describe()

	//co.base.describe() 测试下这么调用行不行
	fmt.Println("原始方法：co.base.describe()", co.base.describe())

	//定义一个接口 一般er结尾 describer() //错误写法  接口后不跟小括号()
	//定义一个接口describer 里边有个describe方法 返回值为string
	type describer interface {
		describe() string
	}
	/*
		可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。
		因为嵌入了 base ，在这里我们看到 container 也实现了 describer 接口。
	*/
	/*
		此话来自教程翻译 下边是自我理解的大白话
		用方法嵌入结构可以用来将接口实现赋予其他结构。在这里，我们看到容器现在实现了描述符接口，因为它嵌入了base。
	*/
	var d describer = co
	fmt.Println("describer:", d.describe())
	//Here we see that a container now implements the describer interface because it embeds base.
	// 问d 为社么嫩调用 describe()方法
	/*
		1--base 绑定了describe()方法 base可以直接调用
		2--base 嵌入到了container结构体 简介初始化了container的co可以调用describe()方法
			并且co.describe() 和 co.base.describe() 两种方式都行
		3--这里我们发现container实现了describer接口 因为container内嵌了base
			而“恰巧”base绑定的方法正好实现了describer接口
			最终发现co 实现了该接口 ==》 于是d 能调用 describe() 方法
		4--可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。
			base结构体带有方法describe() 该方法实现了describer接口 base内嵌到container结构体 container结构体也实现了describer接口
	*/

}
