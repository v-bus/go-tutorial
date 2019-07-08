package main

import (
	"fmt"
	"sync"
	"trydefer/rectarea"
)

func doer(s string) {
	fmt.Println("Hello")
	defer fmt.Println("Buy..")
	if len(s) > 0 {
		switch s {
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

	name := "Bushmin"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: \n")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

	var wg sync.WaitGroup

	r1 := rectarea.NewRect(-67, 89)
	r2 := rectarea.NewRect(5, -67)
	r3 := rectarea.NewRect(8, 9)
	rects := []rectarea.Rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.Area(&wg)
	}
	wg.Wait()
    fmt.Println("All go routines finished executing")
}
