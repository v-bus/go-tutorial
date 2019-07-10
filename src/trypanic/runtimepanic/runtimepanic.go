package runtimepanic

import (
	"fmt"
	"runtime/debug"
)

//A runs runtime error
func A() {
	fmt.Println("Inside {}runtimepanic.A()")
	defer recovery()
	n := []int{1, 2, 3}
	fmt.Println(n[3])
	fmt.Println("normally returned from A()")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovery runtimerror: ", r)
		debug.PrintStack()
	}
}
