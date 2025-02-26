package main

import (
	"fmt"
)

// Employee struct with fields ID, Name, and Salary
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// UpdateSalary updates the salary of an employee
func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

// Display prints the employee details
func (e Employee) Display() {
	fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
}

// AddEmployee adds an employee to the map
func AddEmployee(emp *Employee, db map[int]*Employee) {
	db[emp.ID] = emp
	fmt.Println("Employee added successfully!")
}

// RemoveEmployee removes an employee from the map by ID
func RemoveEmployee(id int, db map[int]*Employee) {
	if _, exists := db[id]; exists {
		delete(db, id)
		fmt.Println("Employee removed successfully!")
	} else {
		fmt.Println("Employee not found!")
	}
}

func main() {
	// Employee database
	employees := make(map[int]*Employee)

	// Adding employees
	emp1 := &Employee{ID: 101, Name: "Alice", Salary: 50000}
	emp2 := &Employee{ID: 102, Name: "Bob", Salary: 60000}
	AddEmployee(emp1, employees)
	AddEmployee(emp2, employees)

	// Display all employees
	fmt.Println("\nEmployee List:")
	for _, emp := range employees {
		emp.Display()
	}

	// Updating salary
	fmt.Println("\nUpdating Alice's Salary to 55000...")
	emp1.UpdateSalary(55000)
	emp1.Display()

	// Removing an employee
	fmt.Println("\nRemoving Bob from system...")
	RemoveEmployee(102, employees)

	// Display remaining employees
	fmt.Println("\nUpdated Employee List:")
	for _, emp := range employees {
		emp.Display()
	}
}
