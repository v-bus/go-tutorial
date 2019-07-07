package main

import (
	"fmt"
	"tryvariadic/changestr"
	"array/basicarray"
)

func main() {
	welcome := []string{"Hello", "World"}
	basicarray.PrintSliceLenCap(welcome)
	changestr.ChangeString(welcome...)
	fmt.Println(welcome)
	basicarray.PrintSliceLenCap(welcome)
	
	welcome = changestr.ChangeString(welcome...)
	fmt.Println(welcome)
	basicarray.PrintSliceLenCap(welcome)
}
