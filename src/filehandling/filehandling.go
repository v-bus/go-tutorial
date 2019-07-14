package main

import (
	"tryerrors/filepatherrors"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File text.txt was not read")
		fmt.Println(filepatherrors.ReturnNames("test.txt"))
		fmt.Println(filepatherrors.ReturnNames("*"))

		return
	}
	fmt.Println("File data: ", string(data))
}
