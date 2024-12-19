// Package library provides a simple implementation for managing a library's book collection.
// It includes functionalities to add, remove, borrow, return, and list books.
package library

import (
	"fmt"
)

// Library represents a collection of books.
type Library struct {
	books map[string]Book // A map of books with their ID as the key.
}

// NewLibrary creates and initializes a new Library instance.
func NewLibrary() *Library {
	return &Library{books: make(map[string]Book)}
}

// AddBook adds a new book to the library.
// If a book with the same ID already exists, it will be overwritten.
func (l *Library) AddBook(book Book) {
	l.books[book.ID] = book
}

// RemoveBook removes a book from the library by its ID.
// If the book does not exist, an error message is displayed.
func (l *Library) RemoveBook(id string) {
	_, exists := l.books[id]
	if !exists {
		fmt.Println("Book not found")
		return
	}
	delete(l.books, id)
}

// BorrowBook marks a book as borrowed by its ID.
// If the book does not exist or is already borrowed, an error message is displayed.
func (l *Library) BorrowBook(id string) {
	book, exists := l.books[id]
	if !exists {
		fmt.Println("Book not found")
		return
	}

	if book.IsBorrowed {
		fmt.Println("Book already is borrowed")
		return
	}
	book.IsBorrowed = true
	l.books[id] = book
}

// ReturnBook marks a book as returned by its ID.
// If the book does not exist or is not borrowed, an error message is displayed.
func (l *Library) ReturnBook(id string) {
	book, exists := l.books[id]
	if !exists {
		fmt.Println("Book not found")
		return
	}

	if !book.IsBorrowed {
		fmt.Println("Book already returned")
		return
	}
	book.IsBorrowed = false
	l.books[id] = book
}

// ListBooks displays all books in the library along with their details.
// It also prints the total count of books in the library.
func (l *Library) ListBooks() {
	var count int
	for key := range l.books {
		fmt.Printf("\nBook ID: %v\n", l.books[key].ID)
		fmt.Printf("Book Title: %v\n", l.books[key].Title)
		fmt.Printf("Book Author: %v\n", l.books[key].Author)
		fmt.Printf("Book IsBorrowed: %v\n", l.books[key].IsBorrowed)
		count++
	}
	fmt.Println("\nTotal books: ", count)
}
