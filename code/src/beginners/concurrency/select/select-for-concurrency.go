// using select to handle channels

package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	//anonlymous functions
	go func() {
		for {
			channel1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			channel2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	//infinite loop
	for {
		select {
		case msg1 := <-channel1:
			// these are inherently blocking calls
			fmt.Println(msg1)

		case msg2 := <-channel2:
			fmt.Println(msg2)
		}

	}

}
