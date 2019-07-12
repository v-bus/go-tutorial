//Package usertypefunc  creates a new function type add which accepts two integer arguments and returns a integer
package usertypefunc

import "fmt"

type add func(a int, b int) int

//DefUseTypedFunc defines variable of type add
func DefUseTypedFunc() {
	var a add = func(a int, b int) int {
		return (a + b)
	}
	s := a(5, 6)
	fmt.Printf("s := a(5, 6) is %d, a() has type %T", s, a)
}
