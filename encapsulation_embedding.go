package main

import "fmt"

// Encapsulation: Private fields (lowercase) and public methods (uppercase)
type BankAccount2 struct {
	accountNumber string // private: only accessible within package
	balance       float64
	owner         string
}

// Constructor function (public)
func NewBankAccount2(accountNumber, owner string, initialBalance float64) *BankAccount2 {
	return &BankAccount2{
		accountNumber: accountNumber,
		owner:         owner,
		balance:       initialBalance,
	}
}

func (b *BankAccount2) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount2) GetOwner() string {
	return b.owner
}

func (b *BankAccount2) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	b.balance += amount
	return nil
}

func (b *BankAccount2) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > b.balance {
		return fmt.Errorf("insufficient funds")
	}
	b.balance -= amount
	return nil
}

type Address struct {
	Street  string
	City    string
	Country string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

type Customer struct {
	Name    string
	Email   string
	Address // embedded struct - promotes Address fields to Customer
}

// SavingsAccount embeds BankAccount2
type SavingsAccount struct {
	*BankAccount2         // embedded pointer to BankAccount2
	InterestRate  float64 // additional field
}

func (s *SavingsAccount) ApplyInterest() {
	interest := s.GetBalance() * s.InterestRate
	s.Deposit(interest)
}

func main() {
	// Encapsulation example
	account := NewBankAccount2("ACC001", "John Doe", 1000.0)
	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())

	account.Deposit(500)
	fmt.Printf("After deposit: $%.2f\n", account.GetBalance())

	err := account.Withdraw(200)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("After withdrawal: $%.2f\n", account.GetBalance())
	}

	customer := Customer{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Address: Address{
			Street:  "Lagos Island",
			City:    "Lagos",
			Country: "Nigeria",
		},
	}

	// Accessing embedded fields directly
	fmt.Printf("\nCustomer: %s\n", customer.Name)
	fmt.Printf("City: %s\n", customer.City) // Direct access to embedded field
	fmt.Printf("Full Address: %s\n", customer.FullAddress())

	// Embedding with BankAccount2
	savings := &SavingsAccount{
		BankAccount2: NewBankAccount2("SAV001", "Alice Brown", 5000),
		InterestRate: 0.05,
	}

	fmt.Printf("\nSavings Account Owner: %s\n", savings.GetOwner())
	fmt.Printf("Balance before interest: $%.2f\n", savings.GetBalance())
	savings.ApplyInterest()
	fmt.Printf("Balance after 5%% interest: $%.2f\n", savings.GetBalance())
}
