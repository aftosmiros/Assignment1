// Package main demonstrates the use of the employee package to manage employees in a company.
// It creates a company, adds employees, and displays their details.
package main

import (
	"github.com/astanait-university/advanced-programming-1-assignment-1/pkg/employee"
)

func main() {
	// Create a new company instance.
	c := employee.NewCompany()

	// Add a full-time employee.
	c.AddEmployee(employee.FullTimeEmployee{
		ID:     1,
		Name:   "Alice",
		Salary: 5000,
	})

	// Add a part-time employee.
	c.AddEmployee(employee.PartTimeEmployee{
		ID:          2,
		Name:        "Bob",
		HourlyRate:  20,
		HoursWorked: 100,
	})

	// List all employees in the company.
	c.ListEmployees()
}
