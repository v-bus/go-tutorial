package filepatherrors

import (
	"fmt"
	"path/filepath"
	"unicode/utf8"
)

//ReturnNames return the names of all files that matches a patter
func ReturnNames(pattern string) (fileNames []string, err error) {
	if utf8.RuneCountInString(pattern) <= 0 {
		return nil, nil
	}
	files, error := filepath.Glob(pattern)
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println("Bad pattern -->", error)
		return nil, error
	}
	fmt.Println(files)
	return files, nil
}
