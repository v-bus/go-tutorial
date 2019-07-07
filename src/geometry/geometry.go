//geometry.go
package main

import (
	"fmt"
	"geometry/rectangle"
)

func main() {
	fmt.Println("Geometrical shape properties")
	var rectLength, rectWidth = 6.0, 7.0
	fmt.Printf("Area is %.2f", rectangle.Area(rectLength, rectWidth))
	fmt.Printf("Diagonal is %.2f", rectangle.Diagonal(rectLength, rectWidth))
}

func init(){
	fmt.Println("init geom")
}