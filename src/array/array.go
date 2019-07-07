package main

import (
	"array/appender"
	"array/basicarray"
	"array/funcall"
	array "array/makeslice"
	"array/multidimensional"
	"array/trycopy"
	"fmt"
)

//tested array named arrayOfAnimals, contains animals names
var arrayOfAnimals [9]string

func init() {
	arrayOfAnimals = [9]string{"lion", "tiger", "cat", "dog", "sheep", "cow", "bow", "dow", "mow"}
}

func main() {
	sliceOfAnimals := arrayOfAnimals[:]
	sliceOfCats := arrayOfAnimals[2:3]

	basicarray.PrintArrayBefore(arrayOfAnimals[:])

	basicarray.PrintSlice(sliceOfAnimals, "___")
	basicarray.PrintSlice(sliceOfCats, "_k_")

	basicarray.PrintSliceLenCap(sliceOfCats)
	basicarray.PrintArrayAfter(arrayOfAnimals[:])

	sliceOfCats = sliceOfCats[:cap(sliceOfCats)]
	basicarray.PrintSliceLenCap(sliceOfCats)
	basicarray.PrintSlice(sliceOfCats, "_mu_")
	basicarray.PrintArrayAfter(arrayOfAnimals[:])

	fmt.Println("make result ", array.MakeSlice(5, 5))

	basicarray.PrintSliceLenCap(sliceOfAnimals)

	sliceOfAnimals = appender.AppendStringToSlice(sliceOfAnimals, []string{"Trakorator"})

	basicarray.PrintSliceLenCap(sliceOfAnimals)

	var newNilSlice []string

	basicarray.PrintSliceLenCap(newNilSlice)

	newNilSlice = appender.AppendStringToSlice(newNilSlice, sliceOfAnimals)

	basicarray.PrintSliceLenCap(newNilSlice)

	funcall.SliceFuncallChangeArray()

	multidimensional.TryMultiSlice()

	trycopy.TryCopy()
}
