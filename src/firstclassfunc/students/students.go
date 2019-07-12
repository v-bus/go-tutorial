//Package students filters a slice of students based on some criteria
package students

import "fmt"

type student struct {
	firstName,
	lastName,
	grade,
	country string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

//RunStudentFilter runs program which filters a slice of students based on grade == "B" criteria
// and prints result slice of students
func RunStudentFilter() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	resS := filter(s, func(stdnt student) bool {
		if stdnt.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println()
	fmt.Println("slice of students before filter is ", s)

	fmt.Println("filtered slice  of students resS is ", resS)

}
