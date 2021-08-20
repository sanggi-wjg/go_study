package main

import (
	"fmt"
	"go_study/accounts"
)

func main() {
	account := accounts.NewAccount("Jay")
	account.Deposit(20)

	// fmt.Println(account)
	// fmt.Println(account.Balance())
	// err := account.WithDraw(10)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = account.WithDraw(10)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// accounts.TryWithDraw(account, 10)
	// accounts.TryWithDraw(account, 10)
	account.ChangeOwner("New User")
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
