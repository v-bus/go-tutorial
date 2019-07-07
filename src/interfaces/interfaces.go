package main

import (
	"fmt"
	"interfaces/apdescriber"
	"interfaces/calcsal"
	"interfaces/describer"
	"interfaces/empop"
	"interfaces/fndvowel"
	"interfaces/slcalc"
	"interfaces/tester"
	"interfaces/zerointerface"
)

func main() {
	//findvowel example
	name := fndvowel.MyString("Tom Anderson")
	var v fndvowel.VowelFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
	fmt.Println()

	//practical use of an interface
	//CalclateSalary example
	perm1 := calcsal.Permanent{
		EmpID:        1,
		BasicPayment: 12000,
		Pf:           5000,
	}
	perm2 := calcsal.Permanent{
		EmpID:        2,
		BasicPayment: 15000,
		Pf:           2000,
	}
	perm3 := calcsal.Permanent{
		EmpID:        3,
		BasicPayment: 10000,
		Pf:           1000,
	}
	cont1 := calcsal.Contract{
		EmpID:        4,
		BasicPayment: 200000,
	}
	cont2 := calcsal.Contract{
		EmpID:        4,
		BasicPayment: 200000,
	}
	fmt.Println("Total salary of employees is ", calcsal.TotalExpence(perm1, perm2, perm3, cont1, cont2))

	//interface is a tuple of (type, value)
	//lets check it
	var t tester.Tester
	myF64 := tester.MyFloat64(65.6)
	t = myF64
	tester.Describe(t)
	t.Test()
	fmt.Println()

	//empty interface implemented by all types
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)

	//assert interface type
	var sa interface{} = 56
	assert(sa) //prints 56
	// assert(s)  //returns panic: interface conversion: interface {} is string, not int
	fmt.Println()

	//assert interface with checking type
	assertVoK(sa)
	assertVoK(s)

	//assert type of an interface by switch interface.(type)
	findType(sa)
	findType(s)
	findType(99.9)

	//assert cunstom type by switch operator
	pers1 := describer.Person{
		Name: "Peet",
		Age:  45,
	}
	describer.FindTypes("Peet")
	describer.FindTypes(pers1)

	//value receiver and poiner receiver implementation
	pers2 := apdescriber.Person{
		Name: "Ron",
		Age:  32,
	}

	pers3 := apdescriber.Person{
		Name: "Cohn",
		Age:  22,
	}
	adr1 := apdescriber.Address{
		City:    "Moscow",
		Address: "Kremlin",
	}
	var apd1 apdescriber.Describer
	apd1 = pers2
	apd1.Describe()

	apd1 = &pers3
	apd1.Describe()

	var apd2 apdescriber.Describer
	// apd2 = adr1 // cannot use adr1 (type apdescriber.Address) as type apdescriber.Describer in assignment:
	// apdescriber.Address does not implement apdescriber.Describer (Describe method has pointer receiver)
	apd2 = &adr1
	apd2.Describe()

	//multiple interface implementation
	var emp1 slcalc.Employee
	emp1.FirstName = "John"
	emp1.LastName = "Smith"
	emp1.BasicPay = 10000
	emp1.Pf = 1000
	emp1.TotalLeaves = 28
	emp1.LeavesTaken = 15

	var sl slcalc.SalaryCalculator = emp1 //call direct interface
	sl.DisplaySalary()
	var l slcalc.LeaveCalculator = emp1 //call direct interface
	fmt.Println("Leaves left are ", l.CalculateLeavesLeft())

	fmt.Println("Embedded interface")

	//Embeded interface
	var eop empop.EmployeeOperations = emp1 //call embeded interface
	eop.DisplaySalary()
	fmt.Println("Leaves left ", eop.CalculateLeavesLeft())

	//nil Interfaces
	var desc zerointerface.Describer

	if desc == nil {
		fmt.Printf("Interface type zerointerface.Describer %T and value %v", desc, desc)
	}
	//desc.Describe() //panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i interface{}) {
	fmt.Printf("Interface has type %T and value %v", i, i)
	fmt.Println()
}

func assert(i interface{}) {
	s := i.(int)
	fmt.Println(s)
}

func assertVoK(i interface{}) {
	val, ok := i.(int)
	fmt.Println(val, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("Type is string and a value is ", i.(string))
	case int:
		fmt.Println("Type is int and a value is ", i.(int))
	default:
		fmt.Println("Unknown type")
	}
}
