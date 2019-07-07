package tester

import "fmt"

//Tester interface
type Tester interface {
	Test()
}

//MyFloat64 type
type MyFloat64 float64

//Test () method implementation
func (m MyFloat64) Test() {
	fmt.Println(m)
}

//Describe prints Tester type and value
func Describe(t Tester) {
	fmt.Printf("Interface Tester has a type  %T and a value %v ", t, t)
	fmt.Println()
}
