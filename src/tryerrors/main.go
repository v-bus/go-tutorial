package main

import (
	"fmt"
	"os"
	"tryerrors/customerrors"
	"tryerrors/dnserrors"
	"tryerrors/filepatherrors"
	"tryerrors/openfile"
)

func getMultArg(args ...interface{}) []interface{} {
	return args
}
func main() {

	//try os.Open errors
	fmt.Println(string(os.PathListSeparator))
	fmt.Println(string(os.PathSeparator))
	fmt.Println(os.UserHomeDir())
	pS := string(os.PathSeparator) //path separator
	fileName := fmt.Sprint(
		getMultArg(os.UserHomeDir())[0],
		pS, "Documents",
		pS, "go",
		pS, "go-tutorial",
		pS, "README.md")
	fmt.Println(fileName)
	openfile.OpenFile(fileName)

	//try net.DNSErrors
	fmt.Println(getMultArg(dnserrors.LookupDNSHost("google.com"))[0])

	//try filepath.Glob errors
	filepatherrors.ReturnNames("[")

	//try custom errors
	fmt.Println("customerrors.RunCircleAreaFunc(-20.0, 3)")
	customerrors.RunCircleAreaFunc(-20.0, 3)
	fmt.Println("customerrors.RunCircleAreaFunc(-20.0, 2)")
	customerrors.RunCircleAreaFunc(-20.0, 2)
	fmt.Println("customerrors.RunCircleAreaFunc(-20.0, 1)")
	customerrors.RunCircleAreaFunc(-20.0, 1)

	//try custom errors by methods
	fmt.Println("customerrors.RunRectArea(-20.0, -40.0)")
	customerrors.RunRectArea(-20.0, -40.0)
	fmt.Println("customerrors.RunRectArea(-0, 0)")
	customerrors.RunRectArea(-0, 0)
	fmt.Println("customerrors.RunRectArea(20.0, -40.0)")
	customerrors.RunRectArea(20.0, -40.0)
	fmt.Println("customerrors.RunRectArea(-20.0, 40.0)")
	customerrors.RunRectArea(-20.0, 40.0)
	fmt.Println("customerrors.RunRectArea(20.0, 40.0)")
	customerrors.RunRectArea(20.0, 40.0)
}
