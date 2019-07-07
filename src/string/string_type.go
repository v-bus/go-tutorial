package main

import "fmt"

func main() {
	firstName := "Viktor"
	secondName := "Bushmin"
	fullName := firstName + " " + secondName
	arrayOfNames := []string{firstName, secondName, fullName}
	fmt.Printf("%q", arrayOfNames)
}