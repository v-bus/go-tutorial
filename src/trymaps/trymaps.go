package main

import (
	"fmt"
	"trymaps/stringintmaptools"
)

func main() {
	personalSalary := map[string]int{
		"steve": 12000,
		"liza":  5000,
		"can":   1000,
	}

	personalSalary["mike"] = 100000

	for key := range personalSalary {
		employee, isEmployee := personalSalary[key]
		if isEmployee {
			fmt.Println("Salary of  ", key, "is ", employee) //, " is ", personalSalary[employee])
		}
	}

	strEmployee := "liza"
	fmt.Println("Length of map is ", len(personalSalary))
	_, isEmployee := personalSalary[strEmployee]
	if isEmployee {
		fmt.Println("Key named ", strEmployee, " has record ", personalSalary[strEmployee], " and will be deleted")
		delete(personalSalary, strEmployee)
	} else {
		fmt.Println("Employee was not found")
	}

	fmt.Println("Length of map has become ", len(personalSalary))

	personalSalary["joeye"] = 33333
	fmt.Println("personalSalary before CopySIMap() ", personalSalary)
	newPersonalSalary := stringintmaptools.CopySIMap(personalSalary)

	// newPersonalSalary["joeye"] = 444444
	fmt.Println("newPersonalSalary after CopySIMap() ", newPersonalSalary)
	fmt.Println("personalSalary after CopySIMap() ", personalSalary)

	if stringintmaptools.EqualSIMap(personalSalary, newPersonalSalary) {
		fmt.Println("personalSalary, newPersonalSalary are equal")
	} else {
		fmt.Println("personalSalary, newPersonalSalary are NOT equal")
	}

}
