package main

import (
	"fmt"
	"math"
)

//Employee struct
type Employee struct {
	name     string
	age      int
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

//Rectangle describes rectangles
type Rectangle struct {
	length, width int
}

//Circle describes circles
type Circle struct {
	radius float64
}

//Area for rectangle
func (r Rectangle) Area() int {
	return r.length * r.width
}

//Area for circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//changeName method trying to  change name in structure Employee
func (e Employee) changeName(newName string) {
	if len(newName) > 0 {
		e.name = newName
	}
}

//changeAge method changes age in structure Employee
func (e *Employee) changeAge(newAge int) {
	e.age = int(math.Abs(float64(newAge)))
}

//Address struct
type Address struct {
	city, address string
}

//Person struct
type Person struct {
	firstName, lastName string
	Address
}

//fullAddress prints city + address in one string
func (a Address) fullAddress() {
	fmt.Println("fullAddress is ", a.city, " ", a.address)
}

func area(r Rectangle) int {
	return r.length * r.width
}

func perimeter(r *Rectangle) int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) perimeter() int {
	return perimeter(r)
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
		age:      24,
	}
	displaySalary(emp1) //Calling displaySalary() function of Employee type
	fmt.Println()
	emp1.displaySalary() //Calling displaySalary() method of Employee type

	//try equal named methods
	r := Rectangle{
		length: 5,
		width:  5,
	}
	c := Circle{radius: 25}
	fmt.Println("Area of rect is ", r.Area())
	fmt.Println("Area of circ is ", c.Area())

	//try poiner recieved value in methods
	fmt.Println("Employee name before change is ", emp1.name)
	emp1.changeName("Bruce")
	fmt.Println("Employee name after change is ", emp1.name)

	fmt.Println("Age of Employee before change is ", emp1.age)
	emp1.changeAge(66)
	fmt.Println("Age of Employee after change is ", emp1.age)

	//pronoted methods
	p := Person{
		firstName: "John",
		lastName:  "Smith",
		Address: Address{
			city:    "Moscow",
			address: "Kremlin",
		},
	}

	p.fullAddress() //accessing fullAddress method of address struct

	//differences between function calling and methods calling
	area(r)
	r.Area()

	rPtr := &r
	// area(rPtr)//cannot use rPtr (type *Rectangle) as type Rectangle in argument to area
	rPtr.Area()

	///difference between pointer recivers and pointer arguments
	perimeter(rPtr)
	rPtr.perimeter()

	// perimeter(r) //cannot use r (type Rectangle) as type *Rectangle in argument to perimeter
	r.perimeter()

	//non-struct type in methods
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is ", sum)
}
