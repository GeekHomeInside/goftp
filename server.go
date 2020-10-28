package main

import (
	"fmt"
	"flag"
	// "os"
	"goftp/utils/server"
)


func main() {
	
	var (
		port string
		username string
		password string
	)
	flag.StringVar(&port, "port", "", "port")
	conn_port := ":" + port
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "username")
	flag.Parse()

	if len(port) == 0 {
			fmt.Println("Please provide a port, username and password\n")
			return
	}
	fmt.Println("Hello I'm GoFTP Server")

	if len(username) > 0 {
		fmt.Println("You have a username:", username)	
	} else {
		fmt.Println("username empty:", username)
	}
	if len(password) > 0 {
		fmt.Println("You have a password:", password)
	} else {
		fmt.Println("password empty:", password)
	}
	server.HandleConnection(conn_port)
	server.Authentification(username,password)
}