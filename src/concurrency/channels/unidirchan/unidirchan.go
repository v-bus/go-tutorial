package unidirchan

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

//RunSendData runs sendData()
func RunSendData() {
	//sendch := make(chan<- int)
	//go sendData(sendch)
	// fmt.Println(<-sendch)//invalid operation: <-sendch (receive from send-only type chan<- int)

	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println("Unidirectional channel data is ", <-chnl)
}
