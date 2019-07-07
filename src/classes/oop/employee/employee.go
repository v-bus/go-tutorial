//Package employee store employee data
package employee

import "fmt"

//employee struct
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

//Employee is exported private struct employee
//type Employee employee

//errorString custom type to store errors for error interface in Employee class
type errorString struct {
	s string
}

//Error implementation error interface method for Employee class
func (err errorString) Error() string {
	return err.s
}

//New func is a constructor of an employee class
func New(firstName, lastName string, totalLeaves, leavesTaken int) employee { //Employee {
	result := employee{firstName, lastName, totalLeaves, leavesTaken}
	return result //Employee(result)
}

//PrintLeavesRemaining prints full name of an employee and calculates and prints difference between Total Leaves and Leaves Taken
func (e employee) PrintLeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
	fmt.Println()
}
