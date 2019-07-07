package trycopy

import (
	"fmt"
)

func init() {
	fmt.Println("package trycopy init()")
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy
}

//TryCopy prints results of copy(dst, src) func by calling countries() func
func TryCopy() {
	fmt.Println(countries())
}
