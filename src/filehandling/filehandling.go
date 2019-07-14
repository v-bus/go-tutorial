package main

import (
	"github.com/gobuffalo/packr/v2"
	"flag"
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
	fmt.Println("File name: ", "test.txt")
	fmt.Println("File data: ", string(data))
}
func getMultArg(args ...interface{}) []interface{} {
	return args
}
func openTestFileAbsoluteFilePath() {
	fileName := fmt.Sprint(
		getMultArg(os.LookupEnv("GOPATH"))[0],
		
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
	fmt.Println("File name: ", fileName)
	fmt.Println("File data: ", string(data))
}

func openTestFileByArgFlag(){
	fptr := flag.String("fpath", "test.txt", "file path to read file")
	flag.Parse()
	fileName := *fptr
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File", fileName, " was not read")
		fmt.Println(filepatherrors.ReturnNames(fileName))
		fmt.Println(filepatherrors.ReturnNames("*"))

		return
	}
	fmt.Println("File name: ", fileName) 
	fmt.Println("File data: ", string(data))
}

func openTestFileByPackrBox(){
	box:=packr.NewBox(fmt.Sprint("..", string(filepath.Separator), "filehandling"))
	data := box.String("test.txt")
	fmt.Println("Content of test.txt:", data)
}
func main() {
	openTestFileByName()
	openTestFileAbsoluteFilePath()
	openTestFileByArgFlag()
	openTestFileByPackrBox()
}
