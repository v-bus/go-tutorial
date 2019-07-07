package main

import "fmt"

func changePointer(val *int) {
	*val = 255
}

func hello() *int {
	i := 5
	return &i
}

func modifyArrayByPtr(arrPtr *[3]int) {
	if arrPtr != nil && len(arrPtr) > 0 {
		//(*arrPtr)[0] = 90 //last sentance
		arrPtr[0] = 90 //shorthand
	}
}

func modifyArrayBySlice(arrPtr []int) {
	if arrPtr != nil && len(arrPtr) > 0 {
		arrPtr[0] = 900
	}
}
func main() {

	//pointer
	b := 255
	var a *int
	a = &b

	fmt.Printf("Type of \"a\" is %T\n", a)
	fmt.Printf("Address of \"b\" is %p", a)
	fmt.Println()

	//nil pointer
	nilA := 5
	var nilB *int
	if nilB == nil {
		fmt.Println("nilB is ", nilB)
		nilB = &nilA
		fmt.Println("nilB after initialization is ", nilB)
	}

	//new() func
	size := new(int)
	fmt.Printf("Type of size is %T, value is %v, address is %p", size, *size, size)
	fmt.Println()
	*size = 85
	fmt.Println("New value of size is ", *size)

	//deferencing pointer
	defB := 255
	defA := &defB

	fmt.Printf("address of defB is %p", defA)
	fmt.Println()
	fmt.Printf("value of defB is %v", *defA)
	fmt.Println()

	*defA++

	fmt.Printf("New value of defB is %v", *defA)
	fmt.Println()

	//passing pointer to func
	*defA = 300

	fmt.Println("Value of devB is ", *defA)

	changePointer(defA)

	fmt.Println("Value of defB after change(*int) func is ", *defA)

	//return pointer from func
	retP := hello()

	fmt.Println("Value of retP returned from hello() func is ", *retP)

	//modify array by pointer
	myArr := [3]int{89, 90, 91}
	modifyArrayByPtr(&myArr)

	fmt.Println("myArr became ", myArr)

	//modify array by slice
	modifyArrayBySlice(myArr[:])

	fmt.Println("myArr became ", myArr)
}
