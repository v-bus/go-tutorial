package appender

import (
	"fmt"
)

func init() {
	fmt.Println("package appender init()")
}

//AppendStringToSlice returns slice that is orgignal slice with appended string
func AppendStringToSlice(slice []string, appeendedString []string) (resultSlice []string) {
	resultSlice = append(slice, appeendedString...)
	fmt.Printf("To slice %v string %v was apennded and result is %v", slice, appeendedString, resultSlice)
	fmt.Println()

	if slice == nil {
		fmt.Println("slice is nil")
	}
	return
}
