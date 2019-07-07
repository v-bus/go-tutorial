package rectangle

import (
		"math"
		"fmt"
)
// Area func
func Area(length, width float64) float64 {
	area := length * width
	return area
}

// Diagonal func
func Diagonal(length, width float64) float64 {
	diagonal := math.Sqrt((length * length) + (width * width))
	return diagonal
}

func init(){
	fmt.Println("rect init")
}
