package slcalc

import (
	"fmt"
	"math"
)

//SalaryCalculator interface
type SalaryCalculator interface {
	DisplaySalary()
}

//LeaveCalculator interface
type LeaveCalculator interface {
	CalculateLeavesLeft() float64
}

// Employee struct
type Employee struct {
	FirstName, LastName      string
	BasicPay, Pf             float64
	TotalLeaves, LeavesTaken float64
}

//DisplaySalary Displays Salary
func (e Employee) DisplaySalary() {
	fmt.Println("", e.FirstName, " ", e.LastName, " has salary", (math.Abs(e.BasicPay) + math.Abs(e.Pf)))
}

//CalculateLeavesLeft float64
func (e Employee) CalculateLeavesLeft() float64 {
	return math.Abs(e.TotalLeaves) - math.Abs(e.LeavesTaken)
}
