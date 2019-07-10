package main

import (
	"fmt"
	"trypanic/abpanic"
	"trypanic/pointpanic"
	"trypanic/runtimepanic"
)

func main() {

	//try nil pointer panic and recovery
	firstName := "Nolan"
	defer fmt.Println("deferred call in main")
	pointpanic.FullName(&firstName, nil)
	fmt.Println("return normally from main")

	//try recovery from another goroutine
	abpanic.A()

	//try runtime error
	runtimepanic.A()
	fmt.Println("return normally from main")
}
