package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func GetText(conn net.Conn) {
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
				fmt.Println("Exiting TCP server cause by:")
				fmt.Println("Could not read data")
				fmt.Println(err)
				return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
				fmt.Println("Exiting TCP server!")
				return
		}
		fmt.Print("-> ", string(netData))
		conn.Write([]byte("goFTP Server received: " + netData))
	}
}