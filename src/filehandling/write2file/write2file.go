//Package write2file describes write file examples
package write2file

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetOutput(os.Stderr)
}

//checkNotEmpty return false when all attributes are zeroed and true when attributes are not zero (empty)
func checkAllAttrNotEmpty(i ...interface{}) bool {
	cntNonEmptyValues := 0

	for _, v := range i {
		if typ := reflect.TypeOf(v); reflect.ValueOf(v) != reflect.Zero(typ) {
			cntNonEmptyValues++
		}
	}
	if cntNonEmptyValues == len(i) {
		return true
	}
	return false

}

//TestCheckAllAttrNotEmpty exports checkAllAttrNotEmpty for testing
func TestCheckAllAttrNotEmpty(t *testing.T, shouldReturn bool, i ...interface{}) {
	if r := checkAllAttrNotEmpty(i); r != shouldReturn {
		t.Errorf("TestCheckAllAttrNotEmpty failed with %v, expected %v, returned %v", i, shouldReturn, r)
	} else {
		t.Log("TestCheckAllAttrNotEmpty is OK")
	}
}

//WriteStringToFile writes incoming string to specified file
func WriteStringToFile(s string, filePath string) {

}
