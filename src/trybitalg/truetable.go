package main

import "fmt"

//https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

func main() {
	for i := 8; i < 16; i++ {
		for j := 1; j < 5; j++ {

			fmt.Printf("%08b XOR %08b = %08b \t| \t", i, j, i^j)
			fmt.Printf("%08b OR %08b = %08b  \t| \t", i, j, i|j)
			fmt.Printf("%08b >> %08b = %08b  \t| \t", i, j, i>>uint32(j))
			fmt.Printf("%08b << %08b = %08b", i, j, i<<uint32(j))
			fmt.Println()
		}
	}
}
