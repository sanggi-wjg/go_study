package main

import (
	"fmt"

	"banking"
)

func main() {
	account := banking.Account{
		owner:   "Jay",
		balance: 100,
	}

	fmt.Println(account)
}
