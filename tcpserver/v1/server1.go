package main

import (
	"fmt"
	// 网络通信包，包含了处理tcp/ip, udp等的方法
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	var listener net.Listener
	var error error
	var conn net.Conn
	mapUsers = make(map[string]int)

	fmt.Println("Starting the server...")
	listener, error = net.Listen("tcp", "localhost:50000")
	checkError(error)
	// 无限循环
	for {
		conn, error = listener.Accept()
		checkError(error)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var error error
	for {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)
		checkError(error)
		input :=  string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}
		ix := strings.Index(input, "says")
		clName := input[0: ix - 1];
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: %v\n", string(buf))
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}