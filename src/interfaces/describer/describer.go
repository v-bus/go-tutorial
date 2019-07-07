package describer

import "fmt"

//Describer interface
type Describer interface {
	Describe()
}

//Person struct
type Person struct {
	Name string
	Age  int
}

//Describe method implementation
func (p Person) Describe() {
	fmt.Println("Name is ", p.Name, ", age is ", p.Age)
}

//FindTypes func
func FindTypes(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Println("Unknow type")
	}
}
