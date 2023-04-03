package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/*Go 标准库的 net/http 包为 HTTP 客户端和服务端提供了出色的支持。 在这个例子中，我们将使用它发送简单的 HTTP 请求。*/

func main() {
	/*向服务端发送一个 HTTP GET 请求。
	http.Get 是创建 http.Client 对象并调用其 Get 方法的快捷方式。
	它使用了 http.DefaultClient 对象及其默认设置。*/
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//打印http response 状态
	fmt.Println("Response status : ", resp.Status)

	//打印response body的前5行内容
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

/*PS D:\GolandProjects\gobyexample> go run .\http-clients.go
Response status :  200 OK
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Go by Example</title>
*/