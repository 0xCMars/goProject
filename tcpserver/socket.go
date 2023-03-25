package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.baidu.com"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	conn, err := net.Dial("tcp", remote);
	if (err != nil) {
		fmt.Printf("Dial connection fail")
		return
	}
	io.WriteString(conn, msg);
	for read {
		count, err = conn.Read(data)
		fmt.Printf("Count is %v\n", count)
		// 判断是否继续读
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	conn.Close();
}