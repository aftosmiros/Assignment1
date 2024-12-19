// Package employee provides functionality for managing employees within a company.
// It includes structures for full-time and part-time employees and methods to manage them.
package employee

import (
	"fmt"
)

// Company represents a company that manages a collection of employees.
// The employees are stored in a map where the key is the employee's ID (as a string),
// and the value is the Employee interface.
type Company struct {
	employees map[string]Employee // A map of employee ID to Employee interface.
}

// NewCompany creates and initializes a new Company instance.
// Returns:
//   - A pointer to the newly created Company.
func NewCompany() *Company {
	return &Company{employees: make(map[string]Employee)}
}

// AddEmployee adds an employee to the company's collection.
// Parameters:
//   - emp: An Employee interface representing the employee to add.
//
// Behavior:
//   - Supports adding FullTimeEmployee and PartTimeEmployee types.
//   - Converts the employee's ID (uint64) to a string to use as the map key.
//   - If the employee type is unsupported, prints an error message.
func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.employees[fmt.Sprintf("%d", e.ID)] = emp
	case PartTimeEmployee:
		c.employees[fmt.Sprintf("%d", e.ID)] = emp
	default:
		fmt.Println("Invalid employee type")
	}
}

// ListEmployees prints the details of all employees in the company.
// It iterates over the company's employee map and calls the GetDetails() method
// on each employee to retrieve their formatted details.
// For each employee, it prints:
//   - The employee's ID (as a string).
//   - The output of the employee's GetDetails() method.
func (c *Company) ListEmployees() {
	for id, emp := range c.employees {
		fmt.Printf("ID: %v -> %v\n", id, emp.GetDetails())
	}
}
