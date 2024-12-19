// Package shapes provides data structures and methods for working with geometric shapes.
// It includes functionality to calculate areas and perimeters for various shapes.
package shapes

import "math"

// Circle represents a geometric circle with a given radius.
type Circle struct {
	Radius float64 // The radius of the circle.
}

// Area calculates the area of the circle using the formula: π * radius².
// Returns:
//   - A float64 value representing the area of the circle.
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Perimeter calculates the perimeter (circumference) of the circle using the formula: 2 * π * radius.
// Returns:
//   - A float64 value representing the perimeter of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}
