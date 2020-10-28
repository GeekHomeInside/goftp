package server

import (
	"fmt"
)

func Authentification(username string, password string) {
	if len(username) < 0 { 
		fmt.Println("enter username")
	 }
	if len(password) < 0 { 
		fmt.Println("enter password")
	 }
	if  len(username) > 0 && len(password) > 0 {
		fmt.Println("You're Authenticate")

	}	
}