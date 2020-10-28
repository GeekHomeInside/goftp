package server

import (
	"fmt"
	"net"
)

func HandleConnection(port string) {
	fmt.Println("Ready to handle connexion")
	ln, err := net.Listen("tcp", port)

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
	GetText(conn)
}