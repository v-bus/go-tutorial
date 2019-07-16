package write2file_test

import (
	"filehandling/write2file"
	"testing"
)

func TestCheckAllAttrNotEmpty(t *testing.T) {
	t.Run("1st",
		func(t *testing.T) {
			write2file.TestCheckAllAttrNotEmpty(t, false, "", "a", 0)
		})
	t.Run("2nd",
		func(t *testing.T) {
			write2file.TestCheckAllAttrNotEmpty(t, true, "parameter", 1, 2.2, true)
		})
}
