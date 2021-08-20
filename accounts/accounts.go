package accounts

import (
	"errors"
	"fmt"
	"log"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// Error
var errNoMoney = errors.New("you can't withdraw")

// Constructor
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Get Method
func (a *Account) Balance() int {
	return a.balance
}

// Get Method
func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	return fmt.Sprintf("[Account]\nowner: %s\nbalance: %d", a.owner, a.balance)
}

// Set Method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Set Method
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// 한번 해봄
func TryWithDraw(a *Account, amount int) {
	err := a.WithDraw(amount)
	if err != nil {
		log.Fatalln(err)
	}
}

// Set Method
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

/* package main

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
*/
