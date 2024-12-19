// Package main provides a command-line interface (CLI) for managing a bank account system.
// It demonstrates the use of the bank package for deposit, withdrawal, checking balance,
// and processing batch transactions.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aftosmiros/Assignment1/pkg/bank"
)

func main() {
	// Initialize a bank account with default values.
	account := bank.BankAccount{
		AccountNumber: "123456789", // Unique account number.
		HolderName:    "John Doe",  // Name of the account holder.
		Balance:       0,           // Initial balance set to zero.
	}

	// Start the CLI loop for user interaction.
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display the menu options to the user.
		fmt.Println("\nBank Account System:")
		fmt.Println("1. Deposit")                      // Option to deposit money into the account.
		fmt.Println("2. Withdraw")                     // Option to withdraw money from the account.
		fmt.Println("3. Get balance")                  // Option to display the current account balance.
		fmt.Println("4. Process transactions (batch)") // Option to process a batch of transactions.
		fmt.Println("5. Exit")                         // Option to exit the application.
		fmt.Print("Choose an option: ")

		// Read the user's choice.
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			// Handle deposit operation.
			fmt.Print("Enter amount to deposit: ")
			scanner.Scan()
			amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil {
				fmt.Println("Invalid amount. Please enter a valid number.")
				continue
			}
			account.Deposit(amount)
		case "2":
			// Handle withdrawal operation.
			fmt.Print("Enter amount to withdraw: ")
			scanner.Scan()
			amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil {
				fmt.Println("Invalid amount. Please enter a valid number.")
				continue
			}
			account.Withdraw(amount)
		case "3":
			// Display the current balance of the account.
			account.GetBalance()
		case "4":
			// Handle batch transaction processing.
			fmt.Print("Enter transactions separated by commas (e.g., 100,-50,200,-30): ")
			scanner.Scan()
			transactionInput := strings.TrimSpace(scanner.Text())
			transactionStrings := strings.Split(transactionInput, ",")
			var transactions []float64

			// Parse transaction strings into float64 values.
			for _, ts := range transactionStrings {
				t, err := strconv.ParseFloat(strings.TrimSpace(ts), 64)
				if err != nil {
					fmt.Printf("Invalid transaction: %s. Skipping...\n", ts)
					continue
				}
				transactions = append(transactions, t)
			}

			// Process the batch of transactions.
			bank.Transaction(&account, transactions)
		case "5":
			// Exit the program gracefully.
			fmt.Println("Exiting...")
			return
		default:
			// Handle invalid menu options.
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
