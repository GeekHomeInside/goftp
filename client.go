package main

import (
	"os"
	"fmt"
	"src/client"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}
	CONN_CONFIG := arguments[1]

	fmt.Println("Hello I'm goClient")
	
	client.handleConnection(CONN_CONFIG)
}