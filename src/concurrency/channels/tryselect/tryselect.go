package main

import (
	"fmt"
	"time"
)

//looking to select operator (switch for channels)
func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func runSelectServers() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

//looking to default value in select
func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process done"
}

func runProcessToReachDefault() {
	output := make(chan string)

	go process(output)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-output:
			fmt.Println("Recieved message -->", v)
			return
		default:
			fmt.Println("Nothing was recieved")
		}
	}
}

func main() {
	runSelectServers()
	runProcessToReachDefault()

	//select protects code from panic all goroutines are asleep
	// ch := make(chan int, 0)

	// select {
	// case <-ch:
	// default:
	// }
}
