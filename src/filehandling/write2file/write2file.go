//Package write2file describes write file examples
package write2file

import (
	"log"
	"os"
	"reflect"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetOutput(os.Stderr)
}

//checkNotEmpty return false when all attributes are zeroed and true when attributes are not zero (empty)
func checkAllAttrNotEmpty(i ...interface{}) bool {
	cntNonEmptyValues := 0

	for _, v := range i {
		if typ := reflect.TypeOf(v);  reflect.ValueOf(v).Interface() != reflect.Zero(typ).Interface() {
			cntNonEmptyValues++
		}
	}
	if cntNonEmptyValues == len(i) {
		return true
	}
	return false

}

//WriteStringToFile writes incoming string to specified file
func WriteStringToFile(s string, filePath string) {

}
