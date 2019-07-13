package query

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

//createQuery prints "Type" and "Value" of incoming interface
func createQuery(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

}

//recover made to protect from reflect.Field(int) "index out of range" panic and "not a struct" panic
func recoverFieldCalls() {
	if r := recover(); r != nil {
		fmt.Println("recover panic from ", r)
	}
}

//createQueryPrintFields prints each field of incoming interface with its position in struct, type and value
func createQueryPrintFields(i interface{}) {
	defer recoverFieldCalls()
	if v := reflect.ValueOf(i); v.Kind() == reflect.Struct {
		fmt.Println("Number of fields ", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field %d has type %T and value %v", i, v.Field(i), v.Field(i))
			fmt.Println()
		}
	}
}
//createQueryPrintQuery prints query concatenated of struct fields
//incoming interface should be struct
func createQueryPrintQuery(i interface{}) {
	defer recoverFieldCalls()
	if t := reflect.TypeOf(i); t.Kind() == reflect.Struct {
		query := fmt.Sprintf("insert into %s values(", t.Name())
		v := reflect.ValueOf(i)
		for cnt := 0; cnt < v.NumField(); cnt++ {
			switch v.Field(cnt).Kind() {
			case reflect.Int:
				if cnt == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(cnt).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(cnt).Int())
				}
			case reflect.String:
				if cnt == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(cnt).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(cnt).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Not a struct")
}

//PrintAllQueries runs createQuery() func which prints "Type", "Value" and "Kind" of incoming interface
// and runs createQueryPrintFields() func which prints each field of incoming interface with its position in struct, type and value
// and runs createQueryPrintQuery() func which prints query concatenated of struct fields
func PrintAllQueries() {
	o := order{
		ordID:      5,
		customerID: 6,
	}

	e := employee{
		name: "Mikhail Gorhgn",
		id: 5,
		address: "Moscow, Kremlin",
		country: "RU",
		salary: 5000,
	}
	createQuery(o)
	createQueryPrintFields(o)
	createQueryPrintQuery(o)

	createQuery(e)
	createQueryPrintFields(e)
	createQueryPrintQuery(e)
}
