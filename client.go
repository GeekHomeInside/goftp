package main

import (
	"os"
	"fmt"
	"goftp/utils/client"
)

const (
	USER = "admin"
	PASSWORD = "admin"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port")
		os.Exit(1)
	}
	CONN_CONFIG := arguments[1]
	CONN_USER := USER
	CONN_PASS := PASSWORD
	client.Authentification(CONN_USER, CONN_PASS)
	client.HandleConnection(CONN_CONFIG)
}