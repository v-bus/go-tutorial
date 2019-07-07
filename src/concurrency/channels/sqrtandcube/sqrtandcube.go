package sqrtandcube

import (
	"fmt"
	"math"
)

//digitPower returns -1 if error
//digit := number % 10
//for power>0 {
//     result = digit*result
//		power--
//}
//
func digitPowerSum(number int, power int) int {
	resultSum := -1
	if number > 0 && power > 0 {
		resultSum = 0
		for number > 0 {
			digit := number % 10
			resultSum += int(math.Pow(float64(digit), float64(power)))
			number /= 10
		}
	}
	return resultSum
}

func sqrt(number int, power int, op chan int) {
	op <- digitPowerSum(number, power)
}

func cube(number int, power int, op chan int) {
	op <- digitPowerSum(number, power)
}

//SqrtAndCube returns smth
func SqrtAndCube() {
	sqrtch := make(chan int)
	cubech := make(chan int)

	number := 589

	go sqrt(number, 2, sqrtch)
	go cube(number, 3, cubech)

	squares, cubes := <-sqrtch, <-cubech
	fmt.Println("sqrt is ", squares)
	fmt.Println("cubes are ", cubes)
	fmt.Println("Final output sqrt+cube of digits ", number, " is ", (squares + cubes))
}



func digit(number int, digch chan int) {
	fmt.Println("Entered to sqrtandcube->digit()")
	for number != 0 {
		digit := number % 10
		fmt.Println("digit ", digit)
		digch <- digit
		number /= 10
	}
	close(digch)

}

func calcPower(number int, power int, chnl chan int) {
	fmt.Println("Entered to calcPower()")
	sum := 0
	digch := make(chan int)
	go digit(number, digch)
	for digit := range digch {
		sum += int(math.Pow(float64(digit), float64(power)))
		fmt.Println("power sum", sum)
	}

	chnl <- sum
}

//RunCalcPower runs calcPower and print sum
func RunCalcPower() {
	number := 589
	sqrtop := make(chan int)
	cubeop := make(chan int)

	go calcPower(number, 3, cubeop)
	go calcPower(number, 2, sqrtop)
	cubes := <-cubeop
	squares := <-sqrtop

	fmt.Println("Final output ", (squares + cubes))
}

//RunDigit runs digit
func RunDigit() {
	chnl := make(chan int)
	go digit(123456789, chnl)
	for v := range chnl {
		fmt.Println("digit", v)
	}
}
