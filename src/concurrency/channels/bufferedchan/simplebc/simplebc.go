package simplebc

import (
	"fmt"
	"sync"
	"time"
)

//Run func runs run
func Run() {
	chnl := make(chan string, 2)
	chnl <- "first"
	chnl <- "second"
	fmt.Println(<-chnl)
	fmt.Println(<-chnl)
}

//RunDeadlock runs code wich make goroutines deadlock
func RunDeadlock() {
	ch := make(chan string, 2)
	ch <- "naveen"
	fmt.Println("ch <- \"naveen\" ")
	ch <- "paul"
	fmt.Println("ch <- \"paul\" ")
	fmt.Println(<-ch) //it does not make deadlock, when someone read channel
	fmt.Println("fmt.Println(<-ch)")
	ch <- "steve"
	fmt.Println("ch <- \"steve\" ")
	fmt.Println(<-ch)
	fmt.Println("fmt.Println(<-ch)")
	fmt.Println(<-ch)
	fmt.Println("fmt.Println(<-ch)")
}
func write(chnl chan int) {
	for i := 0; i < 5; i++ {
		chnl <- i
		fmt.Println("Wrote ", i, "to channel")
	}
	close(chnl)
}

//RunWrite runs write()
func RunWrite() {
	chnl := make(chan int, 2)
	go write(chnl)
	time.Sleep(2 * time.Second)
	for v := range chnl {
		fmt.Println("Read ", v, " from channel")
		time.Sleep(2 * time.Second)
	}
}

//RunLenCap runs len(ch) and cap(ch)
func RunLenCap() {
	chnl := make(chan int, 3)
	chnl <- 1
	chnl <- 2
	fmt.Println("Capacity of channel is ", cap(chnl))
	fmt.Println("Length  of channel is ", len(chnl))
	fmt.Println("READ value from channel", <-chnl)
	fmt.Println("New length of channel is ", len(chnl))
}

//waitGroup code

//waitGroup process code
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Starting process Goroutin ", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine ", i, " ended")
	wg.Done()
}

//RunWaitGroupProcess run process func in waitGroup code block
func RunWaitGroupProcess() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i <= no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished excecuting")
}
