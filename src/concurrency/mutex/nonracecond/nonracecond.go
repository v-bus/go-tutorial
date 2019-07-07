package nonracecond

import (
	"fmt"
	"sync"
)

var x = 0

//It is boring to write two similar run functions
//Use struct as excecuted func attribute

type incrAttr struct {
	wg *sync.WaitGroup
	mx *sync.Mutex
	ch chan bool
}

//Increment func increase global variable x by 1
func Increment(input incrAttr) {
	switch {
	case input.wg != nil:
		fallthrough
	case input.mx != nil:
		input.mx.Lock()
		x = x + 1
		input.mx.Unlock()
		input.wg.Done()
	}
}

//RunMxIncrements runs increment() func "numIncrements" times to try mutex
func RunMxIncrements(numIncrements int, execFunc func(input incrAttr)) {
	var wg sync.WaitGroup
	var mx sync.Mutex
	var input incrAttr
	input.mx = &mx
	input.wg = &wg
	input.ch = make(chan bool, 1) //should be buffered, otherwise panic with all goroutines asleep
	for i := 0; i < numIncrements; i++ {
		input.wg.Add(1)
		go execFunc(input)
	}

	input.wg.Wait()
	fmt.Println("Final condition of x = ", x)
}

//IncrementChanneledMx func increment global variable x to 1 and lock its changes with channel reading
func IncrementChanneledMx(input incrAttr) {
	if input.wg != nil {
		if input.ch != nil {
			input.ch <- true
			x = x + 1
			<-input.ch
		}
		input.wg.Done()
	}
}
