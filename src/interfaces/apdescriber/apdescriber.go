package apdescriber

import "fmt"

//Describer interface
type Describer interface {
	Describe()
}

//Address struct
type Address struct {
	City, Address string
}

//Describe method implementation
func (a *Address) Describe() {
	fmt.Println("Address city is ", a.City, ", address is ", a.Address)
}

//Person struct
type Person struct {
	Name string
	Age  int
}

//Describe method implementation
func (p Person) Describe() {
	fmt.Println("Name is ", p.Name, ", an age is ", p.Age)
}
