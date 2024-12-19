// Package shapes provides data structures and methods for working with geometric shapes.
// It includes an interface for shapes and utilities to calculate and display their properties.
package shapes

import "fmt"

// Shape defines a common interface for geometric shapes.
// Any shape implementing this interface must provide methods to calculate its area and perimeter.
type Shape interface {
	Area() float64      // Calculates the area of the shape.
	Perimeter() float64 // Calculates the perimeter of the shape.
}

// PrintShapeDetails takes a Shape interface and prints its area and perimeter.
// Parameters:
//   - s: A Shape interface representing the geometric shape.
// Behavior:
//   - Calls the Area() and Perimeter() methods of the shape.
//   - Prints the results formatted to two decimal places.
func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
