package multidimensional

import "fmt"

func init(){
	fmt.Println("package multidimensional init()")
}

//TryMultiSlice prints output from multidimensional slice of array
func TryMultiSlice() {
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf(" %s ", v2)
		}
		fmt.Println()
	}
}
