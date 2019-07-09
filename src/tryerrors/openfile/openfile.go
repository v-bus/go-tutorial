package openfile

import (
	"fmt"
	"os"
	"unicode/utf8"
)

//ErrorString stores OpenFile package errors
type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}

//OpenFile opens file or returns error
func OpenFile(fileName string) (f *os.File, e error) {
	f = nil
	e = nil
	if utf8.RuneCountInString(fileName) <= 0 {
		e = error(ErrorString("No file name detected!"))
		return nil, e
	}
	var err error
	f, err = os.Open(fileName)
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path ", err.Path, " failed to open")
		return f, err
	}
	fmt.Println(f.Name(), " was opened successfully")

	return f, nil
}
