package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func deposit(account *BankAccount, amount float64) {
	account.Balance += amount
	fmt.Printf("Deposited $%.2f\n", amount)
}

func withdraw(account *BankAccount, amount float64) {
	if account.Balance >= amount {
		account.Balance -= amount
		fmt.Printf("Withdrew $%.2f\n", amount)
	} else {
		fmt.Println("Insufficient funds!")
	}
}

func main1() {
	myAccount := BankAccount{Owner: "John", Balance: 100.00}
	fmt.Printf("Starting balance: $%.2f\n", myAccount.Balance)

	deposit(&myAccount, 50.00)  //Same account!
	withdraw(&myAccount, 30.00) // Same account

	fmt.Printf("Final balance: $%.2f\n", myAccount.Balance)
}
