package funcall

import "fmt"

func init(){
	fmt.Println("package funcall init()")
}

func changeArray(array []int) {
	for i := range array {
		array[i] -= 2
	}
}

//SliceFuncallChangeArray prints changes of some predefined array my calling func changeArray
func SliceFuncallChangeArray() {
	numbers := []int{8, 9, 12}
	fmt.Println("numbers before calling func changeArray is ", numbers)
	changeArray(numbers)
	fmt.Println("numbers after calling func changeArray is ", numbers)
}
