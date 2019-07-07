package main

import (
	"concurrency/channels/bufferedchan/simplebc"
	"fmt"
)

func main() {
	//try sync.waitGroup
	fmt.Printf("\n----->try sync.waitGroup\n\n\n")
	simplebc.RunWaitGroupProcess()
	//len and cap of channels
	fmt.Printf("\n----->len and cap of channels\n\n\n")
	simplebc.RunLenCap()
	//try buffered channels
	fmt.Printf("\n----->try buffered channels\n\n\n")
	simplebc.Run()

	//try goroutine with buffered channels
	fmt.Printf("\n----->try goroutine with buffered channels\n\n\n")
	simplebc.RunWrite()

	//try imitate deadlock with buffered channels
	fmt.Printf("\n----->try imitate deadlock with buffered channels\n\n\n")
	simplebc.RunDeadlock()

}
