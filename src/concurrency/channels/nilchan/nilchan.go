package nilchan

import "fmt"

// CheckChanAndMake cheks if chan of type int is nil and make it
func CheckChanAndMake() {
	var a chan int
	if a == nil {
		fmt.Println("channel is nil")
		a = make(chan int)
		fmt.Printf("Type of channel is %T", a)
		if a == nil {
			fmt.Println("channel is nil again")
		}
	}

}
