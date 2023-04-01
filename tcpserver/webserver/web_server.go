package main
import (
	"fmt"
	"log"
	"net/http"
)

// 网页服务器发送一个 http.Response 响应，它是通过 http.ResponseWriter 对象输出的
// http.Request描述了客户端请求
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Hello Server handler")
	fmt.Println(req.URL)
	fmt.Fprintf(w, "Hello," + req.URL.Path[1:])
	// io.WriteString(w, "hello, world!\n")
}

func main() {
	// http.HandleFunc 注册了一个处理函数（这里是 HelloServer()）来处理对应 / 的请求。
	// 第一个参数是请求的路径，第二个参数是当路径被请求时，需要调用的处理函数的引用。
	http.HandleFunc("/", HelloServer)
	// 与net.Listen的tcp服务器类似
	err := http.ListenAndServe("10.0.16.16:8545", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}