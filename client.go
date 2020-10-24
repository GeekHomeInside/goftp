package main

import (
	"os"
	f "fmt"
	"net"
	"bufio"
)

func main() {
	f.Println("Hello Client")
	conn, err := net.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
		f.Println(err)
		os.Exit(1)
	}
	f.Println(conn, "GET / HTTP/1.0\r\n")
	_, err = bufio.NewReader(conn).ReadString('\n')
}
