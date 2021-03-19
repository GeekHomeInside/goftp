package server

import (
	"fmt"
	"net"
)

func HandleConnection(port string) {
	fmt.Println("Hello I'm GoFTP Server")
	fmt.Println("Ready to handle connexion")
	ln, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("Exiting TCP server cause by:")
		fmt.Println("Listen not ready")
		fmt.Println(err)
		return
	}

	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Exiting TCP server cause by:")
		fmt.Println("TCP not Accepted")
		fmt.Println(err)
		return
	}
	fmt.Println("TCP Accepted")
	fmt.Println(conn)
	GetText(conn)
}