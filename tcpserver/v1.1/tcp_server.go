package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	flag.Parse()
	// if len(os.Args) != 3 {
	// 	panic("usage: host port")
	// }	
	// hostAndPort := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
	// fmt.Printf("%s\n", hostAndPort)
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	// 格式化成字符串
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	fmt.Printf("%s\n", hostAndPort)
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	// 解析string到ip:port
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '" + hostAndPort + "'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "Listen TCP: ")
	println("Litening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	//
	connFrom := conn.RemoteAddr().String()
	println("Connnection from: ", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead + 1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
			// continue
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
 	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func checkError(err error, msg string) {
	if err != nil {
		panic("ERROR: " + msg + " " + err.Error())
	}
}