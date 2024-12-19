// Package shapes provides data structures and methods for working with geometric shapes.
// It includes functionality to calculate areas and perimeters for various shapes.
package shapes

// Square represents a geometric square with a specified side length.
type Square struct {
	Length float64 // The length of one side of the square.
}

// Area calculates the area of the square using the formula: length².
// Returns:
//   - A float64 value representing the area of the square.
func (s Square) Area() float64 {
	return s.Length * s.Length
}

// Perimeter calculates the perimeter of the square using the formula: 4 * length.
// Returns:
//   - A float64 value representing the perimeter of the square.
func (s Square) Perimeter() float64 {
	return s.Length * 4
}
