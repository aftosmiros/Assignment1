// Package shapes provides data structures and methods for working with geometric shapes.
// It includes functionality to calculate areas and perimeters for various shapes.
package shapes

import "math"

// Triangle represents a geometric triangle defined by the lengths of its three sides.
type Triangle struct {
	SideA float64 // Length of the first side of the triangle.
	SideB float64 // Length of the second side of the triangle.
	SideC float64 // Length of the third side of the triangle.
}

// Area calculates the area of the triangle using Heron's formula:
//   Area = √(s * (s - a) * (s - b) * (s - c))
// where `s` is the semi-perimeter: s = (a + b + c) / 2.
//
// Returns:
//   - A float64 value representing the area of the triangle.
func (t Triangle) Area() float64 {
	halfPerimeter := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(halfPerimeter * (halfPerimeter - t.SideA) * (halfPerimeter - t.SideB) * (halfPerimeter - t.SideC))
}

// Perimeter calculates the perimeter of the triangle using the formula:
//   Perimeter = a + b + c
//
// Returns:
//   - A float64 value representing the perimeter of the triangle.
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
