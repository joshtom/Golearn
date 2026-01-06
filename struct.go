package main

import "fmt"

type Account struct {
	AccountNumber string
	Owner         string
	Balance       float64
	AccountType   string
}

// Deposit adds money to the account
func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	a.Balance += amount
	fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, a.Balance)
}

// Withdraw removes money from the account
func (a *Account) Withdraw(amount float64) bool {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return false
	}
	if amount > a.Balance {
		fmt.Println("Insufficient funds")
		return false
	}
	a.Balance -= amount
	fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, a.Balance)
	return true
}

// GetBalance returns the current balance
func (a Account) GetBalance() float64 {
	return a.Balance
}

// DisplayINfo shows account details
func (a Account) DisplayInfo() {
	fmt.Printf("Account Number: %s\nOwner: %s\nAccount Type: %s\nBalance: $%.2f\n",
		a.AccountNumber, a.Owner, a.AccountType, a.Balance)
}

func structFunc() {
	account := Account{
		AccountNumber: "123456789",
		Owner:         "John Doe",
		Balance:       1000.0,
		AccountType:   "Savings",
	}

	account.DisplayInfo()

	account.Deposit(500)
	account.Withdraw(200)
	account.Withdraw(2000)

	fmt.Println()
	account.DisplayInfo()
}
