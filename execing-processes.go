package main

import (
	"os"
	"os/exec"
	"syscall"
)

/*在前面的例子中，我们了解了生成外部进程的知识，
当我们需要在运行的 Go 流程中访问的外部流程时，便可以执行此操作。
但是有时候，我们只想用其它（也许是非 Go）的进程，来完全替代当前的 Go 进程。
这时，我们可以使用经典的 exec 函数的 Go 的实现。*/

func main() {
	//在这个例子中，我们将执行 ls 命令。
	//Go 要求我们提供想要执行的可执行文件的绝对路径， 所以我们将使用 exec.LookPath 找到它（应该是 /bin/ls）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	//Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。
	//我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	//Exec 同样需要使用环境变量。 这里我们仅提供当前的环境变量。
	env := os.Environ()

	/*这里是真正的 syscall.Exec 调用。
	如果这个调用成功，那么我们的进程将在这里结束，
	并被 /bin/ls -a -l -h 进程代替。
	如果存在错误，那么我们将会得到一个返回值。*/
	exeErr := syscall.Exec(binary, args, env)
	if exeErr != nil {
		panic(exeErr)
	}
	/*注意 Go 没有提供 Unix 经典的 fork 函数。
	一般来说，这没有问题，因为启动协程、生成进程和执行进程，
	已经涵盖了 fork 的大多数使用场景。*/
}

//当我们运行程序时，它会替换为 ls。
//实验验证Windows不支持
/*$ go run execing-processes.go
panic: not supported by windows

goroutine 1 [running]:
main.main()
        D:/GolandProjects/gobyexample/execing-processes.go:35 +0x58
exit status 2

*/
