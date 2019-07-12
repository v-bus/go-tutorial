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

//Every closure is bound to its own surrounding variable
func appendString() func(string) string {
	t := "Hello"
	c := func(s string) string {
		t = t + " " + s
		return t
	}
	return c
}

//RunClosureAppendStr runs closure functions example
//appendString which returns func(string) string
//this resulted func is called 4 times
func RunClosureAppendStr() {
	a := appendString()
	b := appendString()

	fmt.Println(a("World"))
	fmt.Println(b("Everybody"))

	fmt.Println(a("Chase"))
	fmt.Println(b("!"))
}
