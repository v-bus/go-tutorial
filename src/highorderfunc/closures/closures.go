//Package closures demonstrate closures functions in golang
// Closures are a special case of anonymous functions.
// Closures are anonymous functions which access the variables defined outside the body of the function.
package closures

import "fmt"

//SimpleClosureExample prints 5
func SimpleClosureExample() {
	a := 5
	func() {
		fmt.Println("SimpleClosureExample() --> a := ", a)
	}()
}
