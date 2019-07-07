package calcsal

import "fmt"

//SalaryCalculator interface to implement salary operations
type SalaryCalculator interface {
	CalculatorSalary() int
	getEmpID() int
}

//Permanent employees salary struct
type Permanent struct {
	EmpID        int
	BasicPayment int
	Pf           int
}

//Contract employees salary struct
type Contract struct {
	EmpID        int
	BasicPayment int
}

//CalculatorSalary salary of permanent employees method
func (p Permanent) CalculatorSalary() int {
	return p.BasicPayment + p.Pf
}

//CalculatorSalary - salary of contracted employees method
func (c Contract) CalculatorSalary() int {
	return c.BasicPayment
}

func (p Permanent) getEmpID() int {
	return p.EmpID
}

func (c Contract) getEmpID() int {
	return c.EmpID
}

//TotalExpence returns sum of employees salary -1 if error occurs
//EmpID should not be dublicated
func TotalExpence(s ...SalaryCalculator) int {
	expence := 0
	var EmpIDs map[int]int
	EmpIDs = make(map[int]int)
	for _, v := range s {
		_, was := EmpIDs[v.getEmpID()]
		if was == false {
			expence += v.CalculatorSalary()
			if EmpIDs != nil {
				EmpIDs[v.getEmpID()] = 1
			} else {
				fmt.Println("calcsal.go TotalExpence() error \"EmpIDs map is nil\" ")
				expence = -1
				break
			}
		} else {
			fmt.Println("calcsal.go TotalExpence() warning \"dublicated EmpID\" EmpID = ", v.getEmpID())
		}
	}
	return expence
}
