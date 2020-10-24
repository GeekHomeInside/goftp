package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
)

const (
    CONN_HOST = "127.0.0.1"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)


func handleConnection() {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Receive from: ",conn.RemoteAddr() ,status)
}

func main() {
	fmt.Println("Hello I'm goClient")
	handleConnection()
}