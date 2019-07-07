package main

import "fmt"

func doer(s string) {
	fmt.Println("Hello")
	defer fmt.Println("Buy..")
	if len(s) > 0 {
		switch s{
		case "first":
			fmt.Println("First")
		case "second":
			fmt.Println("Second")
		case "third":
			fmt.Println("Third")
		default:
			fmt.Println(s)
		}
	}
}

func main() {
	doer("")
	doer("firsst")
	doer("first")
	doer("second")
	doer("third")
}
