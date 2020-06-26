package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Start server...")

	fd, _ := net.Listen("tcp", ":8081")

	con, _ := fd.Accept()

	for {
		message, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Println("Message received: ", string(message))

		newmessage := strings.ToUpper(message)
		con.Write([]byte(newmessage + "\n"))
	}
}
