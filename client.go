package main

import (
	"flag"
	"fmt"
	"goftp/utils/client"
	"os"
)

func main() {
	var (
		username string
		password string
	)
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "username")
	flag.Parse()
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port")
		os.Exit(1)
	}
	CONN_CONFIG := arguments[1]
	// client.Authentification(username, password)
	client.HandleConnection(CONN_CONFIG)
}