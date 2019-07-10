//Package pointpanic is a simple program to print the full name of a person
package pointpanic

import "fmt"
func recoverName() {
	if r := recover(); r!=nil {
		fmt.Println("recovered from ", r)
	}
}
//FullName prints full name
func FullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("return normally from FullName func")
}
