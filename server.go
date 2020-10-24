package main

import (
	"fmt"
	"net"
)

const (
    CONN_HOST = "127.0.0.1"
    CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)


func handleConnection(conn net.Conn) {
	fmt.Println("Connexion OK\n")
	fmt.Fprintf(conn, "Hello Caller!, I'm GoFTP Server :)")
	fmt.Println("recived from:\n", conn.RemoteAddr())
	conn.Close()
  }

func main() {
	fmt.Println("Hello I'm GoFTP Server")
	ln, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}