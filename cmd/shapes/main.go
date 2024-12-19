// Package main demonstrates the usage of the shapes package by creating
// various geometric shapes and displaying their area and perimeter.
package main

import (
	"github.com/aftosmiros/Assignment1/pkg/shapes"
)

func main() {
	// Create a slice of Shape interface to store different geometric shapes.
	shapeslist := []shapes.Shape{
		shapes.Rectangle{Length: 10, Width: 5},        // Rectangle with length 10 and width 5
		shapes.Circle{Radius: 7},                      // Circle with radius 7
		shapes.Square{Length: 4},                      // Square with side length 4
		shapes.Triangle{SideA: 3, SideB: 4, SideC: 5}, // Triangle with sides 3, 4, and 5
	}

	// Iterate over the slice and print the area and perimeter of each shape.
	for _, shape := range shapeslist {
		shapes.PrintShapeDetails(shape)
	}
}
