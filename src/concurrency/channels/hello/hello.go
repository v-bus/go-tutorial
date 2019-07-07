package hello

import (
	"fmt"
	"time"
)

//hello func prints "Hello!"
func hello(done chan bool) {
	fmt.Println("Hello!")
	done <- true
}

//RunHello runs hello() func
func RunHello() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("leavine hello runner")
}

func helloDelaied(done chan bool) {
	fmt.Println("helloDelaied func going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("helloDelaied func awake and going to write do done channel")
	done <- true
}


//RunHelloDelaied calls helloDelaied and returns true to done channel
func RunHelloDelaied() {
	done := make(chan bool)
	fmt.Println("RunHelloDetailed() going to call helloDetailed func")
	go helloDelaied(done)
	<-done
	fmt.Println("RunHelloDelaied() recieved data")
}
