package fndvowel

import "strings"

//VowelFinder interface declaration
type VowelFinder interface {
	FindVowels() []rune
}

//MyString type
type MyString string

//FindVowels methods returns []rune of vowels in ms string
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	s := string(ms)
	ms = MyString(strings.ToLower(s))
	for _, rune := range ms {
		switch rune {
		case 'a', 'i', 'e', 'o', 'y', 'u':
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
