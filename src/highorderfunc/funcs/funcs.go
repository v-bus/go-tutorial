package funcs

import "fmt"

func simple(a func(a int, b int) int) {
	fmt.Println("a(60, 7) is ", a(60, 7))
}

func simpleReturnFunc() (a func(a int, b int) int) {
	a = func(a int, b int) int {
		return a + b
	}
	return
}

//RunHighOrderFuncAttr runs high-ordered function simple()
func RunHighOrderFuncAttr() {
	//simple func is high-ordered because it applies func as incomong attribute
	f := func(a int, b int) int {
		return a + b
	}

	simple(f)
}

//RunHighOrderReturnFunc runs high-ordered function simpleReturnFunc() which returns function (a, b int) int
func RunHighOrderReturnFunc() {
	//simple func is high-ordered because it returns func as result
	s := simpleReturnFunc()
	fmt.Println("s := simpleReturnFunc() is s(60,7)", s(60, 7))
}
