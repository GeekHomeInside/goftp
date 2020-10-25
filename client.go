package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
	"strings"
)


func handleConnection(CONN_CONFIG string) {
	conn, err := net.Dial("tcp", CONN_CONFIG)
	if err != nil {
		fmt.Println("Exiting TCP server cause by:")
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
				fmt.Println("TCP client exiting...")
				return
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}
	CONN_CONFIG := arguments[1]

	fmt.Println("Hello I'm goClient")
	
	handleConnection(CONN_CONFIG)

}