package racecond

import (
	"fmt"
	"sync"
)

//Program with race condition
var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}


//RunIncrements runs "increment()" func "numIncrements" times
func RunIncrements(numIncrements int) {
	var wg sync.WaitGroup
	for i := 0; i < numIncrements; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final value is ", x)
}
