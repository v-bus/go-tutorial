//Package main calls examples of firstclass functions tutorial
package main

import (
	"firstclassfunc/students"
	"firstclassfunc/usertypefunc"
	"firstclassfunc/anonfunc"
)

func main() {
	anonfunc.AssignFuncToVariable()
	anonfunc.CallAnonFunc()
	anonfunc.PassArgToAnonFunc()

	usertypefunc.DefUseTypedFunc()

	students.RunStudentFilter()
}
