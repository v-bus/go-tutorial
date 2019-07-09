package customerrors

import (
	"errors"
	"fmt"
	"math"
)

//simple program which calculates the area of a circle and will return an error if the radius is negative

func circleArea(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, errors.New("Area calculation failed, radius is less or equal zero")
	}
	return math.Pi * math.Pow(radius, 2.0), nil
}

//get errors by formatted text
func circleAreaErrorf(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less or equal zero", radius)
	}
	return math.Pi * math.Pow(radius, 2.0), nil
}

func circleAreaError(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, &areaError{"Area calculation failed, radius is less or equal zero", radius}
	}
	return math.Pi * math.Pow(radius, 2.0), nil
}

//RunCircleAreaFunc runs circleArea() func
func RunCircleAreaFunc(radius float64, fcase int) {

	switch fcase {
	case 1:
		area, err := circleArea(radius)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Area of the circle with radius %0.2f is %0.2f", radius, area)
			fmt.Println()
		}
	case 2:
		area, err := circleAreaErrorf(radius)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Area of the circle with radius %0.2f is %0.2f", radius, area)
		fmt.Println()
	case 3:
		area, err := circleAreaError(radius)
		if err != nil {
			if err, ok := err.(*areaError); ok {
				fmt.Printf("Radius %0.2f is less than zero\n", err.radius)
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("Area of the circle with radius %0.2f is %0.2f", radius, area)
		fmt.Println()
	default:
		fmt.Println("No circleArea() func was chosen")
	}
}

// now try return error struct
type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

//try error methods
type rectAreaError struct {
	err    string
	width  float64
	length float64
}

func (e *rectAreaError) Error() string {
	return fmt.Sprintf("width %0.2f and length %0.2f: %s", e.width, e.length, e.err)
}

func (e *rectAreaError) widthNegative() bool {
	return e.width < 0
}

func (e *rectAreaError) lengthNegative() bool {
	return e.length < 0
}

func rectArea(width, length float64) (float64, error) {
	err := ""
	if width < 0 {
		err += "width is less than zero "
	}
	if length < 0 {
		if err == "" {
			err += "length is less than zero"
		} else {
			err += ", length is less than zero"
		}
	}
	if err != "" {
		return 0, &rectAreaError{err, width, length}
	}
	return (length * width), nil
}

//RunRectArea runs rectArea() func
func RunRectArea(width, length float64) {
	area, err := rectArea(width, length)
	if err != nil {
		if err, ok := err.(*rectAreaError); ok {
			if err.lengthNegative() {
				fmt.Printf("length %0.2f is negative", err.length)
				fmt.Println()
			}
			if err.widthNegative() {
				fmt.Printf("width %0.2f is negative", err.width)
				fmt.Println()
			}
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle with width=%0.2f and length=%0.2f is %0.2f", width, length, area)
	fmt.Println()
}
