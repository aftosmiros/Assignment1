// Package employee provides structures and interfaces for managing employees in a company.
// It supports both full-time and part-time employees and provides methods to retrieve employee details.
package employee

import "fmt"

// Employee defines a common interface for all employee types.
// Each employee type must implement the GetDetails method, which returns a formatted string of the employee's details.
type Employee interface {
	GetDetails() string // Returns the details of the employee as a formatted string.
}

// FullTimeEmployee represents an employee who works full-time.
// Fields:
//   - ID: Unique identifier for the employee.
//   - Name: Name of the employee.
//   - Salary: Annual salary of the employee.
type FullTimeEmployee struct {
	ID     uint64 // Unique identifier for the full-time employee.
	Name   string // Name of the full-time employee.
	Salary uint32 // Annual salary of the full-time employee.
}

// GetDetails returns a formatted string with the details of the full-time employee.
// It includes the employee's ID, name, and annual salary.
func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-Time Employee\nID: %v\nName: %v\nSalary: %v\n", f.ID, f.Name, f.Salary)
}

// PartTimeEmployee represents an employee who works part-time.
// Fields:
//   - ID: Unique identifier for the employee.
//   - Name: Name of the employee.
//   - HourlyRate: The rate the employee is paid per hour.
//   - HoursWorked: The total number of hours the employee has worked.
type PartTimeEmployee struct {
	ID          uint64  // Unique identifier for the part-time employee.
	Name        string  // Name of the part-time employee.
	HourlyRate  uint64  // Hourly rate for the part-time employee.
	HoursWorked float32 // Total hours worked by the part-time employee.
}

// GetDetails returns a formatted string with the details of the part-time employee.
// It includes the employee's ID, name, hourly rate, and total hours worked.
func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part-Time Employee\nID: %v\nName: %v\nHourly Rate: %v\nHours Worked: %.2f\n",
		p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}
