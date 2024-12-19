// Package services provides business logic to process transactions
// and manage operations on bank accounts using the models package.
package bank

import (
	"fmt"
)

// Transaction processes a batch of transactions for a given bank account.
// Each transaction in the slice can be:
//   - Positive: Represents a deposit to the account.
//   - Negative: Represents a withdrawal from the account.
//   - Zero: Skipped, as no operation is required.
//
// Parameters:
//   - account: A pointer to the BankAccount struct representing the account to operate on.
//   - transactions: A slice of float64 values representing deposit or withdrawal amounts.
//
// Behavior:
//   - Positive values in the transactions slice call the Deposit method on the account.
//   - Negative values call the Withdraw method (absolute value is used).
//   - Zero values are skipped with a warning message.
func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			// Handle deposit
			account.Deposit(amount)
		} else if amount < 0 {
			// Handle withdrawal (convert negative value to positive for Withdraw method)
			account.Withdraw(-amount)
		} else {
			// Skip zero-value transactions
			fmt.Println("Skipping transaction with amount 0.")
		}
	}
}
