package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var conn net.Conn
	var err error
	var inputReader *bufio.Reader
	var input string
	var clientName string

	conn, err = net.Dial("tcp", "localhost:50000")
	checkError(err);

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ = inputReader.ReadString('\n');
	trimmedClient := strings.Trim(clientName, "\n") // 去掉\n
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ = inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		} 
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		checkError(err)
	}
}
// if there is a tiny error, the whole server have to stop
func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}