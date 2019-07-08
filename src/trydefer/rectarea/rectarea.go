package rectarea

import (
	"fmt"
	"sync"
)

//Rect struct represents rectangle with and length
type Rect struct {
	length int //shold not be less than zero
	width  int //shold not be less than zero
	area   int //shold not be less than zero
}

//New creates zero Rect object
//with length = 0, width = 0 and area = -1
func New() *Rect {
	return &Rect{0, 0, -1}
}

//NewRect creates Rect object with given "length" and "width" attributes
//if incoming attributes are less than zero they become zero
func NewRect(length, width int) Rect {
	if length < 0 {
		length = 0
	}
	if width < 0 {
		width = 0
	}
	area := -1
	return Rect{length, width, area}
}

//PrintArea func print "area" elenent of "Rect" structure
func (r Rect) PrintArea() {
	fmt.Printf("rect %v's area %d", r, r.area)
	fmt.Println()
}

//Area func calculates area of rect
func (r Rect) Area(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println()
	if r.width <= 0 {
		fmt.Printf("%v's width should be greater than zero\n", r)
		return
	}
	if r.length <= 0 {
		fmt.Printf("%v's length should be greater than zero\n", r)
		return
	}

	r.area = r.length * r.width
	r.PrintArea()

}
