package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"tryerrors/filepatherrors"
)

func openTestFileByName() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File text.txt was not read")
		fmt.Println(filepatherrors.ReturnNames("test.txt"))
		fmt.Println(filepatherrors.ReturnNames("*"))

		return
	}
	fmt.Println("File data: ", string(data))
}
func getMultArg(args ...interface{}) []interface{} {
	return args
}
func openTestFileAbsoluteFilePath() {
	fileName := fmt.Sprint(
		getMultArg(os.LookupEnv("GOPATH"))[0],
		string(filepath.Separator),
		"src",
		string(filepath.Separator),
		"filehandling",
		string(filepath.Separator),
		"test.txt")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File", fileName, " was not read")
		fmt.Println(filepatherrors.ReturnNames(fileName))
		fmt.Println(filepatherrors.ReturnNames("*"))

		return
	}
	fmt.Println("File data: ", string(data))
}
func main() {
	openTestFileByName()
	openTestFileAbsoluteFilePath()
}
