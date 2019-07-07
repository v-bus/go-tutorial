package empop

import "interfaces/slcalc"

// EmployeeOperations interface
type EmployeeOperations interface {
	slcalc.LeaveCalculator
	slcalc.SalaryCalculator
}