// Package main provides a command-line interface (CLI) for managing a library's book collection.
// It demonstrates the use of the library package to add, borrow, return, and list books.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/astanait-university/advanced-programming-1-assignment-1/pkg/library"
)

func main() {
	// Initialize a new library instance.
	// This creates an empty library ready to manage books.
	lib := library.NewLibrary()

	// Create a scanner to read user input from the command line.
	// Scanner is used for taking multi-line input from the user.
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display the menu options to the user.
		fmt.Println("\nLibrary Management System:")
		fmt.Println("1. Add a book")     // Option to add a new book to the library.
		fmt.Println("2. Borrow a book")  // Option to borrow a book from the library.
		fmt.Println("3. Return a book")  // Option to return a borrowed book.
		fmt.Println("4. List all books") // Option to list all books in the library.
		fmt.Println("5. Exit")           // Option to exit the application.
		fmt.Print("Choose an option: ")

		// Read the user's choice from input.
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text()) // Trim whitespace for cleaner input.

		// Process the user's choice using a switch statement.
		switch choice {
		case "1":
			// Add a new book to the library.
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter book author: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			// Create a new Book instance and add it to the library.
			book := library.Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			lib.AddBook(book)
			fmt.Println("Book added successfully.")

		case "2":
			// Borrow a book from the library using its ID.
			fmt.Print("Enter book ID to borrow: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			lib.BorrowBook(id)

		case "3":
			// Return a borrowed book to the library using its ID.
			fmt.Print("Enter book ID to return: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			lib.ReturnBook(id)

		case "4":
			// List all books currently in the library.
			fmt.Println("Listing all books in the library:")
			lib.ListBooks()

		case "5":
			// Exit the program gracefully.
			fmt.Println("Exiting...")
			return

		default:
			// Handle invalid input from the user.
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
