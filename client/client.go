package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		readMes := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := readMes.ReadString('\n')
		fmt.Fprintf(con, text+"\n")
		message, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
