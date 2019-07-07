package closedchan

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

//RunProducer runs producer()
func RunProducer() {
	chnl := make(chan int)
	go producer(chnl)
	for {
		v, ok := <-chnl
		if ok == false {
			break
		}
		fmt.Println("Recieved", v, ok)
	}
}

//RunProducerRange runs producer calling channel by for v:=range channel
func RunProducerRange() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Recieved ", v)
	}
}
