package main

import (
	"fmt"
	"structures/computer"
)

//Employee structure
type Employee struct {
	firstName, secondName string
	age, salary           int
}

//Person structure
type Person struct {
	string
	int
	Address
}

//Address structure
type Address struct {
	city, address string
}

func main() {
	//creating named structure
	emp1 := Employee{
		firstName:  "Sam",
		secondName: "Smith",
		age:        30,
		salary:     15000,
	}

	emp2 := Employee{"John", "McBay", 25, 25000}

	fmt.Println(emp1)
	fmt.Println(emp2)

	//creating anonymous structure
	emp3 := struct {
		firstName, secondName string
		age, salary           int
	}{
		firstName:  "Albert",
		secondName: "Dohn",
		age:        55,
		salary:     30000,
	}

	fmt.Println(emp3)

	//zero value of a structure
	var emp4 Employee
	fmt.Println(emp4)

	//zero some fields of a structure
	emp5 := Employee{
		firstName:  "Lisa",
		secondName: "Simpson",
	}

	fmt.Println(emp5)

	//accessing individual fields
	emp6 := Employee{"Bard", "Simpson", 13, 2}
	fmt.Println(emp6.firstName)
	fmt.Println(emp6.secondName)
	fmt.Println(emp6.age)
	fmt.Println(emp6.salary)

	//zero struct with added fields
	var emp7 Employee
	emp7.firstName = "Lethery"
	emp7.secondName = "Beethorn"

	fmt.Println(emp7)

	//pointer to a  struct
	emp8Ptr := &Employee{"Lila", "Longer", 0, 0}
	fmt.Printf("Address of a struct is %p", emp8Ptr)
	fmt.Printf(" with value %v", *emp8Ptr)
	fmt.Println()

	fmt.Println("First name is ", emp8Ptr.firstName)
	fmt.Println("Age is ", emp8Ptr.age)

	//anonymous fields
	person := Person{"Jane", 25, Address{"", ""}}
	fmt.Println(person)

	//anonymous fields default names
	person.string = "Nelson"
	person.int = 33
	fmt.Println(person)

	//anonymous nested structure fields
	person.Address = Address{"Moscow", "remlin"}
	fmt.Println("person record is ", person)

	fmt.Println("Name is ", person.string)
	fmt.Println("Age is ", person.int)
	fmt.Println("City is ", person.Address.city)
	fmt.Println("Address is ", person.Address.address)

	//promoted fieldes
	fmt.Println("City is ", person.city)
	fmt.Println("Address is ", person.address)

	mySpec := computer.Spec{
		Maker: "Apple",
		Price: 1000,
	}
	fmt.Println(mySpec)

	//structures equality
	mySpec2 := mySpec
	var mySpec3 computer.Spec
	mySpec3.Maker = "Google"
	mySpec3.Price = 2000
	switch {
	case mySpec == mySpec2:
		fmt.Println("mySpec == mySpec2")
		fallthrough
	case mySpec != mySpec3:
		fmt.Println("mySpec != mySpec3")
	default:
		fmt.Println("mySpec != mySpec2")
		fmt.Println("mySpec == mySpec3")
	}

	//not comparable structures
	// if mySpec == person { //invalid operation: mySpec == person (mismatched types computer.Spec and Person)
	// 	fmt.Println("mySpec==person")
	// }

	
}
