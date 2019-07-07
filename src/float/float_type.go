package main

import "fmt"

func main(){
	a, b := 34.65, 12.67
	fmt.Printf("type of \"a\" is %T, type of b is %T\n",a, b)
	sum := a + b
	diff := a - b
	fmt.Printf("sum is %v, sum type is %T diff is %v, diff type is %T\n", sum, sum, diff, diff)
	
	no1, no2 := 3, 4
	fmt.Printf("no1 is %v, no1 type is %T \nno2 is %v, no2 type is %T, \nno1+no2=%v, no1-no2=%v",no1, no1, no2, no2, no1+no2, no1-no2)
	
}