package main

import (
	"fmt"
	"golwee/banking"
)

func main() {
	account := banking.CreateBankAccount("woodi")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
	account.ChangeOwner("mircat")
	fmt.Println(account)

}
