package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"tryerrors/filepatherrors"

	"github.com/gobuffalo/packr"
)

func openTestFileByName() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File text.txt was not read")
		fmt.Println(filepatherrors.ReturnNames("test.txt"))
		fmt.Println(filepatherrors.ReturnNames("*"))

		return
	}
	fmt.Println("openTestFileByName()")
	fmt.Println("File name: ", "test.txt")
	fmt.Println("File data: ", string(data))
}
func getMultArg(args ...interface{}) []interface{} {
	return args
}
func openTestFileAbsoluteFilePath() {
	fmt.Println("openTestFileAbsoluteFilePath")
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
	fmt.Println("File name: ", fileName)
	fmt.Println("File data: ", string(data))
}

func openTestFileByArgFlag() {
	fmt.Println("openTestFileByArgFlag")
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

func openTestFileByPackrBox() {
	fmt.Println("openTestFileByPackrBox")
	box := packr.NewBox("../filehandling")
	data := box.String("test.txt")
	fmt.Println("Content of test.txt:", data)
}

func openTestFileByBufio3Bytes() {
	fmt.Println("openTestFileByBufio3Bytes")
	var fptr *string //flag string argument pointer
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	if flag.Lookup("fpath") == nil {
		fptr = flag.String("fpath", "test.txt", "open test.txt file by 3 bytes")
	} else {
		fstr := flag.Lookup("fpath").Value.String()
		fptr = &fstr
		fmt.Println(*fptr)
	}
	flag.Parse() //do after flags parameters set

	fh, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fh.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(fh)
	
	for  {
		b := make([]byte, 3)
		_, err = r.Read(b)
		if err!=nil{
			fmt.Println("Error reading file:",err)
			break
		}
		fmt.Println(string(b))
	}
}
func main() {
	openTestFileByName()
	openTestFileAbsoluteFilePath()
	openTestFileByArgFlag()
	openTestFileByPackrBox()
	openTestFileByBufio3Bytes()
}
