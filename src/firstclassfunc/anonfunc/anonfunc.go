//Package anonfunc demonstrate anonymous golang func creation and using
//functions are called anonymous functions since they do not have a name
//The only way to call this function is using the variable
package anonfunc

import "fmt"

//AssignFuncToVariable assigns a function to a variable
func AssignFuncToVariable() {
	a := func() {
		fmt.Println("Assigned func() fmt.Println()")
	}
	a()
	fmt.Printf("\"a\"() has a type %T ", a)
}

//CallAnonFunc calls a anonymous function without assigning it to a variable
func CallAnonFunc() {
	func() {
		fmt.Println()
		fmt.Println("Call anonymous function with func() fmt.Println() call")
	}()
}

//PassArgToAnonFunc passes arguments to anonymous functions just like any other function
func PassArgToAnonFunc(){
		func(n string) {
		fmt.Println()
		fmt.Println("Call anonymous function func(n string) where \"n\" is ", n)
	}("Hello string")
}