package main

import (
	"fmt"
	"os"
)

/*使用 os.Exit 可以立即以给定的状态退出程序*/

func main() {
	//当使用 os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用。
	defer fmt.Println("!!!")
	os.Exit(3) //退出并且退出状态为 3。
	//exit status 3
}

/*注意，不像例如 C 语言，Go 不使用在 main 中返回一个整数来指明退出状态。
如果你想以非零状态退出，那么你就要使用 os.Exit。*/
/*
如果你使用 go run 来运行 exit.go，那么退出状态将会被 go 捕获并打印。

$ go run exit.go
exit status 3
通过编译并执行一个二进制文件的方式，你可以在终端中查看退出状态。

$ go build exit.go
$ ./exit
$ echo $?
3*/
/*shibinbin@shibinbin00 MINGW64 /d/GolandProjects/gobyexample (master)
$ ./exit.exe

shibinbin@shibinbin00 MINGW64 /d/GolandProjects/gobyexample (master)
$ echo $?
3

*/
//注意，程序中的 ! 永远不会被打印出来。
