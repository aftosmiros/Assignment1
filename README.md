# Multi-System Application

This Go project is the **first assignment** for the **Advanced Programming 1** course at **Astana IT University**. It has been developed by **Tulebayev Miras**. The project contains multiple reusable packages (`pkg`) for different systems, including a **Library Management System**, a **Bank Account System**, a **Shapes Management System**, and an **Employee Management System**. Each package is modular and implements its functionality with clear interfaces and methods, adhering to Go's best practices.

---

## **Project Structure**

```
multi_system_app/
├── cmd/
│   ├── library/
│   │   └── main.go              # CLI for Library Management System
│   ├── bank/
│   │   └── main.go              # CLI for Bank Account System
│   ├── shapes/
│   │   └── main.go              # CLI for Shapes Management System
│   ├── employee/
│   │   └── main.go              # CLI for Employee Management System
├── pkg/
│   ├── library/
│   │   ├── book.go              # Book struct and logic
│   │   ├── library.go           # Library management logic
│   ├── bank/
│   │   ├── account.go           # BankAccount struct and methods
│   │   ├── transaction.go       # Logic for transactions
│   ├── shapes/
│   │   ├── shapes.go            # Shape interface and utility
│   │   ├── rectangle.go         # Rectangle logic
│   │   ├── circle.go            # Circle logic
│   │   ├── square.go            # Square logic
│   │   ├── triangle.go          # Triangle logic
│   ├── employee/
│       ├── employee.go          # Interfaces and structs for employees
│       ├── company.go           # Logic for managing employees in a company
├── go.mod                       # Go module file
```

---

## **Packages Overview**

### 1. **Library Management System** (`pkg/library`)

The `library` package provides functionality to manage a library's book collection. It allows adding, borrowing, returning, and listing books.

#### **Key Features**
- **Add Books**: Add new books to the library.
- **Borrow Books**: Mark books as borrowed.
- **Return Books**: Mark borrowed books as returned.
- **List Books**: Display all books with their availability status.

#### **Usage Example**
```go
lib := library.NewLibrary()
lib.AddBook(library.Book{ID: "101", Title: "Go Programming", Author: "John Doe"})
lib.BorrowBook("101")
lib.ListBooks()
```

---

### 2. **Bank Account System** (`pkg/bank`)

The `bank` package simulates a bank account system. It supports depositing, withdrawing, checking balance, and batch processing transactions.

#### **Key Features**
- **Deposit Money**: Add money to the account.
- **Withdraw Money**: Deduct money from the account if sufficient funds are available.
- **Check Balance**: Display the current account balance.
- **Batch Transactions**: Process multiple deposits and withdrawals.

#### **Usage Example**
```go
account := bank.BankAccount{AccountNumber: "12345", HolderName: "John Doe", Balance: 0}
account.Deposit(1000)
account.Withdraw(200)
bank.Transaction(&account, []float64{500, -100, 300})
account.GetBalance()
```

---

### 3. **Shapes Management System** (`pkg/shapes`)

The `shapes` package provides functionality to calculate the area and perimeter of various geometric shapes, such as rectangles, circles, squares, and triangles.

#### **Key Features**
- **Shape Interface**: Generic interface for calculating area and perimeter.
- **Shapes**:
  - Rectangle: `shapes.Rectangle`
  - Circle: `shapes.Circle`
  - Square: `shapes.Square`
  - Triangle: `shapes.Triangle`
- **Utility Function**: `PrintShapeDetails` displays the area and perimeter of any shape.

#### **Usage Example**
```go
shapesList := []shapes.Shape{
    shapes.Rectangle{Length: 10, Width: 5},
    shapes.Circle{Radius: 7},
    shapes.Square{Length: 4},
    shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
}
for _, shape := range shapesList {
    shapes.PrintShapeDetails(shape)
}
```

---

### 4. **Employee Management System** (`pkg/employee`)

The `employee` package provides functionality to manage employees in a company, supporting both full-time and part-time employees.

#### **Key Features**
- **Employee Interface**: Shared interface for all employee types.
- **Full-Time Employees**: `employee.FullTimeEmployee` stores salary-based employees.
- **Part-Time Employees**: `employee.PartTimeEmployee` stores hourly-based employees.
- **Company**: `employee.Company` manages employees and provides methods to add and list employees.

#### **Usage Example**
```go
c := employee.NewCompany()
c.AddEmployee(employee.FullTimeEmployee{ID: 1, Name: "Alice", Salary: 5000})
c.AddEmployee(employee.PartTimeEmployee{ID: 2, Name: "Bob", HourlyRate: 20, HoursWorked: 100})
c.ListEmployees()
```

---

## **How to Run the Project**

### Prerequisites
- **Go 1.20+** installed on your system.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/YourUsername/multi_system_app.git
   cd multi_system_app
   ```

2. Run any CLI example:
   - **Library Management**:
     ```bash
     go run cmd/library/main.go
     ```
   - **Bank Account Management**:
     ```bash
     go run cmd/bank/main.go
     ```
   - **Shapes Management**:
     ```bash
     go run cmd/shapes/main.go
     ```
   - **Employee Management**:
     ```bash
     go run cmd/employee/main.go
     ```

---

## **Future Enhancements**

1. **Data Persistence**:  
   Add functionality to save and load data from files or a database.

2. **Web API Integration**:  
   Expose the packages' functionality via RESTful APIs.

3. **Unit Tests**:  
   Add comprehensive test coverage for all packages.

4. **Advanced Features**:
   - Library: Search and filter books.
   - Bank: Add support for multiple accounts.
   - Shapes: Extend with more shapes (e.g., polygons).
   - Employee: Add advanced filters and reports.

---

## **About the Developer**

This project is the **first assignment** for the **Advanced Programming 1** course at **Astana IT University**.  
It was developed by **Tulebayev Miras**.  

- **Email**: [tmiras0706@gmail.com](mailto:tmiras0706@gmail.com)  
- **GitHub**: [github.com/aftosmiros](https://github.com/aftosmiros)

---

## **License**

This project is licensed under the MIT License. See the `LICENSE` file for details.
```

---

### **Changes Added**
1. **Assignment Context**:
   - Mentioned that this project is part of the **Advanced Programming 1** course at **Astana IT University**.
   - Included the name of the developer: **Tulebayev Miras**.

2. **Improved Structure**:
   - Highlighted how each package works with examples.

3. **Professional Tone**:
   - Improved readability and professionalism for academic submission.

Let me know if you need additional tweaks! 🚀