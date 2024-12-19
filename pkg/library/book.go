package library

// Book represents a book in the library.
// It contains the book's ID, title, author, and its borrowing status.
type Book struct {
	ID         string // Unique identifier for the book.
	Title      string // Title of the book.
	Author     string // Author of the book.
	IsBorrowed bool   // Indicates whether the book is currently borrowed.
}

// NewBook creates and initializes a new Book instance.
// Parameters:
//   - id: A unique identifier for the book.
//   - title: The title of the book.
//   - author: The author of the book.
//   - isborrowed: The initial borrowing status of the book.
// Returns:
//   - A pointer to the newly created Book instance.
func NewBook(id string, title string, author string, isborrowed bool) *Book {
	return &Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: isborrowed,
	}
}
