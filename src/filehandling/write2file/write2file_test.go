package write2file_test

import (
	"fmt"
	"testing"

	"github.com/viktor-bushmin/go-tutorial/src/filehandling/write2file"
)

func TestMain(m *testing.M) {

}
func TestWriteStringToFile(t *testing.T) {
	t.Log("Running WriteStringToFile test...")
	t.Run("Exception",
		func(t *testing.T) {
			t.Log("Try to put empty values, should return error")
			err := write2file.WriteStringToFile("", "")
			if err != nil {
				t.Log("With empty input generated error was", err)
			} else {
				t.Error("With empty input should generate error but it did not, returned error nil")
			}
		})
	t.Run("Exception",
		func(t *testing.T) {
			t.Log("Try to put bad filePath")
			err := write2file.WriteStringToFile("", "/home/user1/mypath") //fix test
			fmt.Print(err.Error())
			if err != nil {
				t.Log("With bad filePath  generated error was", err)
			} else {
				t.Error("With bad filePath should generate error but it did not, returned error nil")
			}
		})

}
