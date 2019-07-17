package write2file

import (
	"testing"
)

func TestCheckAllAttrNotEmpty(t *testing.T) {
	t.Run("1st",
		func(t *testing.T) {
			if checkAllAttrNotEmpty("", "a", 0) == false {
				t.Log("checkAllAttrNotEmpty(\"\", \"a\", 0)==false - OK")
			} else {
				t.Error("checkAllAttrNotEmpty(\"\", \"a\", 0)==false - FAILED")
			}
		})
	t.Run("2nd",
		func(t *testing.T) {
			if checkAllAttrNotEmpty("parameter", 1, 2.2, true)== true {
				t.Log("checkAllAttrNotEmpty(\"parameter\", 1, 2.2, true)== true - OK")
			} else {
				t.Error("checkAllAttrNotEmpty(\"parameter\", 1, 2.2, true)== true - FAILED")
			}
		})
		t.Run("3rd",
		func(t *testing.T) {
			if checkAllAttrNotEmpty("","")== false {
				t.Log("checkAllAttrNotEmpty(\"\",\"\")== false - OK")
			} else {
				t.Error("checkAllAttrNotEmpty(\"\",\"\")== false - FAILED")
			}
		})
		t.Run("4th",
		func(t *testing.T) {
			if checkAllAttrNotEmpty("Hello world","C:\\Users\\user1\\Documents\\go\\go-tutorial\\src\filehandling\\write2file\\write2file_internal_test.go")== true {
				t.Log("checkAllAttrNotEmpty(\"Hello world\",\"C:\\Users\\user1\\Documents\\go\\go-tutorial\\src\filehandling\\write2file\\write2file_internal_test.go\")== true - OK")
			} else {
				t.Error("checkAllAttrNotEmpty(\"Hello world\",\"C:\\Users\\user1\\Documents\\go\\go-tutorial\\src\filehandling\\write2file\\write2file_internal_test.go\")== true - FAILED")
			}
		})
}
