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
}
