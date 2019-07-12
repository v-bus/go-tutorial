package main

import (
	"highorderfunc/closures"
	"highorderfunc/funcs"
)

func main() {
	funcs.RunHighOrderFuncAttr()
	funcs.RunHighOrderReturnFunc()

	closures.SimpleClosureExample()
	closures.RunClosureAppendStr()
}
