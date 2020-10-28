package main

import (
	"fmt"
	"os"
	"goftp/utils/server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
			fmt.Println("Please provide port number")
			return
	}
	fmt.Println("Hello I'm GoFTP Server")
	CONN_PORT := ":" + arguments[1]
	server.HandleConnection(CONN_PORT)
}