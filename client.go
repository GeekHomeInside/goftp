package main

import (
	"os"
	"fmt"
	"goftp/utils/client"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port")
		os.Exit(1)
	}
	
	client.HandleConnection(CONN_CONFIG)
}