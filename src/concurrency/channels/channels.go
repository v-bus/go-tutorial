package main

import (
	"concurrency/channels/closedchan"
	"concurrency/channels/hello"
	"concurrency/channels/nilchan"
	"concurrency/channels/sqrtandcube"
	"concurrency/channels/unidirchan"
	"fmt"
)

func breaker(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println(".")
	}
	fmt.Println("----->", s)
}
func main() {

	//chan can be nil and should be initiated by make()
	breaker("chan can be nil and should be initiated by make()")
	nilchan.CheckChanAndMake()

	//try chan in hello func
	breaker("try chan in hello func")
	hello.RunHello()

	//try chan and timer in hello func
	breaker("try chan and timer in hello func")
	hello.RunHelloDelaied()

	// This program will print the sum of the squares and cubes of the individual digits of a number.
	// For example if 123 is the input, then this program will calculate the output as
	// squares = (1 * 1) + (2 * 2) + (3 * 3)
	// cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
	// output = squares + cubes = 50
	breaker("This program will print the sum of the squares and cubes of the individual digits of a number.\nFor example if 123 is the input, then this program will calculate the output as\nsquares = (1 * 1) + (2 * 2) + (3 * 3) \ncubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3) \noutput = squares + cubes = 50")
	sqrtandcube.SqrtAndCube()

	//unidirestional channels
	breaker("Run unidirectional channel.")
	unidirchan.RunSendData()

	//closing channels checkes by v, ok := <- chan
	breaker("Closing channels checkes by v, ok := <- chan")
	closedchan.RunProducer()

	//closing channels for v := range chan
	breaker("closing channels for v := range chan")
	closedchan.RunProducerRange()

	// for res:=range channel used in this example
	breaker("This program will print the sum of the squares and cubes of the individual digits of a number.\nFor example if 123 is the input, then this program will calculate the output as\nsquares = (1 * 1) + (2 * 2) + (3 * 3) \ncubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3) \noutput = squares + cubes = 50")
	sqrtandcube.RunCalcPower()

}
