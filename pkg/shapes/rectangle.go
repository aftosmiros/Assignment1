// Package shapes provides data structures and methods for working with geometric shapes.
// It includes functionality to calculate areas and perimeters for various shapes.
package shapes

// Rectangle represents a geometric rectangle with a specified length and width.
type Rectangle struct {
	Length float64 // The length of the rectangle.
	Width  float64 // The width of the rectangle.
}

// Area calculates the area of the rectangle using the formula: length * width.
// Returns:
//   - A float64 value representing the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Perimeter calculates the perimeter of the rectangle using the formula: 2 * (length + width).
// Returns:
//   - A float64 value representing the perimeter of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Length) * 2
}
