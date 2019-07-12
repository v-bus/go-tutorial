//Package anonfunc demonstrate anonymous golang func creation and using
package anonfunc

import "fmt"

//AssignFuncToVariable assigns a function to a variable
func AssignFuncToVariable() {
	a := func() {
		fmt.Println("Assigned func() fmt.Println()")
	}
	a()
	fmt.Printf("a func() has a type %T ", a)
}
