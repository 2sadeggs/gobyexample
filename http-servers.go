package main

import (
	"fmt"
	"net/http"
)

/*
使用 net/http 包，我们可以轻松实现一个简单的 HTTP 服务器。
*/

//handlers 是 net/http 服务器里面的一个基本概念。
//handler 对象实现了 http.Handler 接口。
//编写 handler 的常见方法是，在具有适当签名的函数上使用 http.HandlerFunc 适配器。
func hello(w http.ResponseWriter, r *http.Request) {
	//handler 函数有两个参数，http.ResponseWriter 和 http.Request。
	//response writer 被用于写入 HTTP 响应数据，这里我们简单的返回 “hello\n”。
	fmt.Fprintf(w, "hello\n")
}

/*这个 handler 稍微复杂一点，
我们需要读取的 HTTP 请求 header 中的所有内容，并将他们输出至 response body。*/
func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v:%v\n", name, h)
		}
	}
}
func main() {
	//使用 http.HandleFunc 函数，可以方便的将我们的 handler 注册到服务器路由。
	//它是 net/http 包中的默认路由，接受一个函数作为参数。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}

/*PS D:\GolandProjects\gobyexample> curl localhost:8080/hello
curl : 无法识别该 URI 前缀。
所在位置 行:1 字符: 1
+ curl localhost:8080/hello
+ ~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotImplemented: (:) [Invoke-WebRequest], NotSupportedException
    + FullyQualifiedErrorId : WebCmdletIEDomNotSupportedException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand

PS D:\GolandProjects\gobyexample> curl http://localhost:8080/hello


StatusCode        : 200
StatusDescription : OK
Content           : hello

RawContent        : HTTP/1.1 200 OK
                    Content-Length: 6
                    Content-Type: text/plain; charset=utf-8
                    Date: Sun, 02 Apr 2023 05:10:25 GMT


Forms             : {}
Headers           : {[Content-Length, 6], [Content-Type, text/plain; charset=utf-8], [Date, Sun, 02 Apr 2023 05:10:25 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 6



PS D:\GolandProjects\gobyexample> curl http://localhost:8080/headers


StatusCode        : 200
StatusDescription : OK
Content           : User-Agent:Mozilla/5.0 (Windows NT; Windows NT 10.0; zh-CN) WindowsPowerShell/5.1.19041.1682

RawContent        : HTTP/1.1 200 OK
                    Content-Length: 93
                    Content-Type: text/plain; charset=utf-8
                    Date: Sun, 02 Apr 2023 05:10:46 GMT

                    User-Agent:Mozilla/5.0 (Windows NT; Windows NT 10.0; zh-CN) WindowsPowerShell/5.1.1...
Forms             : {}
Headers           : {[Content-Length, 93], [Content-Type, text/plain; charset=utf-8], [Date, Sun, 02 Apr 2023 05:10:46 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 93



PS D:\GolandProjects\gobyexample>
*/
