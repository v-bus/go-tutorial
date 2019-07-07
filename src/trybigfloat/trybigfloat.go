package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Float).SetPrec(20).SetInf(false)
	fmt.Printf("%T %f", a, a)
}
