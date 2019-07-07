package intfinder

import "fmt"

func init() {
	fmt.Println("package intfinder init()")
}

//FindInt is variadic func and finds number in numbers and prints index of it
func FindInt(number int, numbers ...int) {
	fmt.Printf("Type of numbers is %T", numbers)
	fmt.Println()

	found := false // flag to switch did we find number
	for iteratorOfNumbers, valueOfNumbers := range numbers {
		if valueOfNumbers == number {
			fmt.Println(number, " was found in index #", iteratorOfNumbers, " in ", numbers)
			found = true
		}
	}

	if !found {
		fmt.Println(number, " was not found in", numbers)
	}
}

//FindIntSliced is based on slice finds number in numbers and print index of it
func FindIntSliced(number int, numbers []int) {
	fmt.Printf("Type of numbers is %T", numbers)
	fmt.Println()

	found := false // flag to switch did we find number
	for iteratorOfNumbers, valueOfNumbers := range numbers {
		if valueOfNumbers == number {
			fmt.Println(number, " was found in index #", iteratorOfNumbers, " in ", numbers)
			found = true
		}
	}

	if !found {
		fmt.Println(number, " was not found in", numbers)
	}
}
