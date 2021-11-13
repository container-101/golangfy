package banking

import "errors"

type BankAccount struct{
	owner string
	balance int
}

// Create New BankAccount
func CreateBankAccount(owner string) *BankAccount{
	account := BankAccount{owner: owner, balance: 0}
	return &account
}
// Deposit x amount on BankAccount
// Modify Actual Receiver
func (a *BankAccount) Deposit(amount int){
	a.balance += amount
}
// Balance of BankAccount
func (a BankAccount) Balance() int{
	return a.balance
}

// Withdraw Balance
func(a *BankAccount) Withdraw(amount int) error{
	result := a.balance - amount
	if result < 0 {
		return errors.New("Can't withdraw(Less than Balance)")
	}
	a.balance = result
	return nil
}