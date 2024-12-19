// Package models provides the core data structures and methods for simulating
// a bank account system. It includes the BankAccount struct and its associated
// methods for managing account operations like deposit, withdraw, and balance retrieval.
package bank

import (
	"fmt"
)

// BankAccount represents a bank account with basic attributes and operations.
type BankAccount struct {
	AccountNumber string  // The unique identifier for the bank account.
	HolderName    string  // The name of the account holder.
	Balance       float64 // The current balance of the account.
}

// Deposit adds a specified amount to the bank account balance.
// Parameters:
//   - amount: The amount to be deposited (must be positive).
//
// Behavior:
//   - If the amount is less than or equal to 0, the deposit is rejected.
//   - Otherwise, the amount is added to the balance and the updated balance is displayed.
func (ba *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	ba.Balance += amount
	fmt.Printf("Deposited: %v, Balance: %v\n", amount, ba.Balance)
}

// Withdraw deducts a specified amount from the bank account balance.
// Parameters:
//   - amount: The amount to be withdrawn (must be positive and <= balance).
//
// Behavior:
//   - If the amount is less than or equal to 0, the withdrawal is rejected.
//   - If the amount exceeds the current balance, the withdrawal is rejected with an error message.
//   - Otherwise, the amount is deducted from the balance and the updated balance is displayed.
func (ba *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return
	}
	if amount > ba.Balance {
		fmt.Println("Insufficient balance")
		return
	}
	ba.Balance -= amount
	fmt.Printf("Withdrew: %v, Balance: %v\n", amount, ba.Balance)
}

// GetBalance prints the details of the bank account, including:
//   - Account Number
//   - Holder Name
//   - Current Balance
func (ba *BankAccount) GetBalance() {
	fmt.Printf("Account Number: %v\nHolder Name: %v\nBalance: %v\n", ba.AccountNumber, ba.HolderName, ba.Balance)
}
