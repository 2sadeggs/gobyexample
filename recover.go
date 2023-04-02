package main

import "fmt"

/*
Go 通过使用 recover 内置函数，可以从 panic 中 恢复recover 。 recover 可以阻止 panic 中止程序，并让它继续执行。
*/
/*在这样的例子中很有用：当其中一个客户端连接出现严重错误，服务器不希望崩溃。
相反，服务器希望关闭该连接并继续为其他的客户端提供服务。
实际上，这就是Go的 net/http 包默认对于 HTTP 服务器的处理。*/

//定义一个panic函数
func mayPanic() {
	panic("a problem")
}

func main() {
	//必须在 defer 函数中调用 recover。 当跳出引发 panic 的函数时，defer 会被激活， 其中的 recover 会捕获 panic。

	//recover 的返回值是在调用 panic 时抛出的错误。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Received. Error:\n", r)
		}
	}()
	//调用自定义panic
	mayPanic()
	//这行代码不会执行，因为 mayPanic 函数会调用 panic。 main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束。
	fmt.Println("After mayPanic()")
}

/*
本文总结2023.04.01
1 panic recover 及defer 用法书写顺序  defer匿名函数中嵌套recover 去捕捉panic
	且必须在defer中调用recover
2 defer在main函数执行完再执行
3 推测 defer 在函数return 后在执行
*/
