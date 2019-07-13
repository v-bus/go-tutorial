package query

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

//prints "Type" and "Value" of incoming interface
func createQuery(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

}

//RunCreateQueryPrintTV run createQuery() func which prints "Type" and "Value" of incoming interface
func RunCreateQueryPrintTV() {
	o := order{
		ordID:      5,
		customerID: 6,
	}
	createQuery(o)
}
