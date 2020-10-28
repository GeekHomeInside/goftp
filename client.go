package main

import (
	"os"
	"fmt"
	"goftp/utils/client"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}
	CONN_CONFIG := arguments[1]

	fmt.Println("Hello I'm goClient")
	
	client.HandleConnection(CONN_CONFIG)
}