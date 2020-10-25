package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
)

func handleConnection(CONN_PORT string) {
	fmt.Println("Ready to handle connexion\n")
	ln, err := net.Listen("tcp", CONN_PORT)

	if err != nil {
		fmt.Println("Exiting TCP server cause by:")
		fmt.Println(err)
		return
	}

	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Exiting TCP server cause by:")
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
				fmt.Println("Exiting TCP server cause by:")
				fmt.Println(err)
				return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
				fmt.Println("Exiting TCP server!")
				return
		}
		fmt.Print("-> ", string(netData))
		conn.Write([]byte("goFTP Server received: " + netData))
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
			fmt.Println("Please provide port number")
			return
	}
	fmt.Println("Hello I'm GoFTP Server")
	CONN_PORT := ":" + arguments[1]
	handleConnection(CONN_PORT)
}