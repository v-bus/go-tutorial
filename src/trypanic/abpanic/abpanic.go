// Package abpanic .
// Recover works only when it is called from the same goroutine.
// It's not possible to recover from a panic that has happened in a different goroutine.
package abpanic

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered: ", r)
	}
}

//A runs goroutine b()
func A() {
	defer recovery()
	fmt.Println("Inside {}abpanic.A()")
	// go b() //recovery() doesnot work
	b() //recovery() works fine
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("inside {}abpanic.b()")
	panic("B paniced")
}
