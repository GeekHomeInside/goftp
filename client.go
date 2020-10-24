package main

import (
	f "fmt"
	"net"
	"bufio"
)

func main() {
	f.Println("Hello Client")
	conn, err := net.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
	}
	f.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
}