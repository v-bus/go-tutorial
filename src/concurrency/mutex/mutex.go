package main

import (
	"concurrency/mutex/nonracecond"
	"concurrency/mutex/racecond"
	"fmt"
)

func main() {
	//launch race conditioned code
	fmt.Println("launch race conditioned code---->")
	racecond.RunIncrements(10000)

	//launch mutexed increment code with sync.Mutex
	fmt.Println("launch mutexed increment code with sync.Mutex------>")
	nonracecond.RunMxIncrements(10000, nonracecond.Increment)


	//launch mutexed increment code with channel
	fmt.Println("launch mutexed increment code with channel------>")
	nonracecond.RunMxIncrements(10000, nonracecond.IncrementChanneledMx)
}
