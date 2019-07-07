package main

import "fmt"

/*
/* func number returns number
*/
func number() int {
	return 75
}
func main(){
	switch num:=number(); {
	case num < 50:
		fmt.Println("num less then 50")
		fallthrough
	case num < 100:
		fmt.Println("num less then 100")
		fallthrough
	case num < 200:
		fmt.Println("num less then 200")
	}
}