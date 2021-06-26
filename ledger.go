/*
	Ledger program that will dynamically build a key value pair for users
*/
package main

import (
	"fmt"
)

func main() {
	//save users and order
	ledger := make(map[string]int)
	usernum := 0
	username := ""
	UserInput(username, usernum, ledger)

}

func UserInput(username string, usernum int, ledger map[string]int) {
	fmt.Println("Welcome to the ledger\nPlease enter your name")
	fmt.Scanln(&username)
	fmt.Println("Please enter your phone number with area code ex: 8505551234")
	fmt.Scanln(&usernum)

	ledger = BuildLedger(ledger, usernum, username)
	for key, value := range ledger {
		fmt.Println(key, value)
	}
	UserInput(username, usernum, ledger)
}

func BuildLedger(ledger map[string]int, usernum int, username string) map[string]int {
	ledger[username] = usernum
	return ledger
}
