package main

import (
	f "fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	f.Fprintf(conn, "Hello Caller!")
	// do more stuff with conn
	conn.Close()
  }

func main() {
	f.Println("Hello Server")
	ln, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		f.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			f.Println(err)
		}
		go handleConnection(conn)
	}
}