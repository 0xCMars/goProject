// using TCP to connect remote sever port 80, and then try using udp protocol, 
// Finnally using tcp to connect IPv6 address
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.0.32.10:80")
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "192.0.32.10:80")
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") // IPv6
	checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
}