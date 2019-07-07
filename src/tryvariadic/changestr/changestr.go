package changestr

import "fmt"

func init() {
	fmt.Println("package changestr init()")
}

//ChangeString changes [0] element to "Go" returns slice []string
func ChangeString(s ...string) []string {
	if cap(s) > 0 {
		s[0] = "Go"
	}
	s = append(s, "playground")
	fmt.Println("Printed from changeStr ", s)
	return s
}
