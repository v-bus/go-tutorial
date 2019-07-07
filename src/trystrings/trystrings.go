package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c  ", s[i])
	}
}

func printCharsRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c  ", runes[i])
	}
}
func printBytesAndChars(s string) {
	printChars(s)
	fmt.Println()
	printBytes(s)
	fmt.Println()
}

func printBytesAndCharsRunes(s string) {
	fmt.Println()
	for index, rune := range s {
		fmt.Printf("%c - %d", rune, index)
		fmt.Println()
	}
}

func printLength(s string) {
	fmt.Printf("Length of %s is %d \n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	if len(s) > 0 {
		s[0] = 'u'
	}
	return string(s)
}

func main() {
	myString := "Hello World"
	fmt.Println(myString)
	printBytesAndChars(myString)

	name := "Señor"
	fmt.Println(name)
	printBytesAndChars(name)

	printCharsRune(name)
	printBytesAndCharsRunes(name)

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // hex bytes
	fmt.Printf("%#v \n", byteSlice)
	byteSliceStr := string(byteSlice)
	fmt.Println(byteSliceStr)

	byteSliceDecimal := []byte{67, 97, 102, 195, 169} // decimal bytes
	fmt.Printf("%#v \n", byteSliceDecimal)
	byteSliceStr = string(byteSliceDecimal)
	fmt.Println(byteSliceStr)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} // runes of unicode
	fmt.Printf("%x \n", runeSlice)
	str := string(runeSlice)
	fmt.Println(str)

	printLength(name)
	fmt.Println("len() of", name, " is ", len(name))
	printLength(byteSliceStr)
	printLength(myString)

	fmt.Println(mutate([]rune(name)))

}
