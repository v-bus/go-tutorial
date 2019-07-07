package basicarray

import "fmt"

func init() {
	fmt.Println("package basicarray init()")
}

/*
/* function prints incoming slice
*/
func PrintSlice(slice []string, suffix string) {
	for i := range slice {
		slice[i] = slice[i] + suffix
	}
	fmt.Println(slice)
}

/*
/* function prints length and capacity of incoming array slice
*/
func PrintSliceLenCap(slice []string) {
	fmt.Printf("array %T length is %v", slice, len(slice))
	fmt.Println()
	fmt.Printf("array %T  capacity is %v", slice, cap(slice))
	fmt.Println()

}

//printArrayBefore prints array
func PrintArrayBefore(array []string) {
	fmt.Println("array before", array)
	fmt.Println()
}

//PrintArrayAfter prints array with words "array after"
func PrintArrayAfter(array []string) {
	fmt.Println("array after", array)
	fmt.Println()
}
