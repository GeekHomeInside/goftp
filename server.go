package main

import (
	f "fmt"
	"net"
)

func main() {
	f.Println("Hello Server")
	ln, err := net.Listen("tcp", "*:1234")
	if err != nil {
		f.Fprintln("Handler Error")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			f.Fprintln("Handler Error")
		}
		go handleConnection(conn)
	}
}