package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*对于操作文件系统中的 目录 ，Go 提供了几个非常有用的函数。*/
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//在当前工作目录下，创建一个子目录。
	//err := os.Mkdir("subdir", 0755)
	//err := os.Mkdir("directories-subdir", 0755) //先注释
	//创建一个已经存在的目录会报错：
	//panic: mkdir directories-subdir: Cannot create a file when that file already exists.
	//check(err)
	//创建这个临时目录后，一个好习惯是：使用 defer 删除这个目录。 os.RemoveAll 会删除整个目录（类似于 rm -rf）。
	//defer os.RemoveAll("directories-subdir") //先临时注释掉 因为下边演示会用到 删除反而看不到效果

	/*一个用于创建临时文件的帮助函数。*/
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("directories-subdir/file1")

	/*我们还可以创建一个有层级的目录，使用 MkdirAll 函数，并包含其父目录。 这个类似于命令 mkdir -p。*/
	err := os.MkdirAll("directories-subdir/parent/child", 0755)
	check(err)

	createEmptyFile("directories-subdir/parent/file2")
	createEmptyFile("directories-subdir/parent/file3")
	createEmptyFile("directories-subdir/parent/child/file4")

	//ReadDir 列出目录的内容，返回一个 os.DirEntry 类型的切片对象。
	c, err := os.ReadDir("directories-subdir/parent")
	check(err)
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	/*Listing subdir/parent
	  child true
	  file2 false
	  file3 false
	*/
	//cd 回到最开始的地方。
	err = os.Chdir("../../..")
	check(err)

	//当然，我们也可以遍历一个目录及其所有子目录。 Walk 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	//Visiting subdir
}

/*filepath.Walk 遍历访问到每一个目录和文件后，都会调用 visit。*/
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
