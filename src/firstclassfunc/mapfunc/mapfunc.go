// Package mapfunc demonstrate map functions
// kind of functions which operate on every element of a collection
package mapfunc

import (
	"fmt"
	"math/rand"
)

func intMap(s []int, f func(int) int) []int {
	var resS []int
	for _, v := range s {
		resS = append(resS, f(v))
	}
	return resS
}

func recoverZeroDivisionPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

//RunIntMapMultiplyTo5 runs intMap with function which multiplies all numbers in slice to 5
func RunIntMapMultiplyTo5() {
	a := []int{1, 2, 3, 4, 45, 5, 6, 55, 7, 8, 9, 0}
	resInt := intMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println("all elements of slice ", a, " was multiplied with intMap func to 5 and slice became ", resInt)

}

//RunIntDivByNom rand.Intn(100) was divided with intMap func by all elements of slice
//zero division panic recovered
func RunIntDivByNom() {

	a := intMap(make([]int, 10), func(n int) int {
		res :=  rand.Intn(100)
		for res <= 0  {
			res = rand.Intn(100)
		}
		return res
	})
	fmt.Println("Randomized randomizer limits produced ", a)
	defer recoverZeroDivisionPanic()
	a = intMap(a, rand.Intn)

	fmt.Println("Sice was generated", a)

	resInt := intMap(a, func(n int) int {
		return rand.Intn(100) / n
	})

	fmt.Println("rand.Intn(100) was divided with intMap func by all elements of slice ", a, " and slice became ", resInt)
}
