// using channels which are like pipes
// to send and receive messages as a blocking operation

package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string)
	go count("sheep", msgChannel)

	// for {
	// 	msg := <-msgChannel
	// 	fmt.Println(msg)
	// }

	// //channel can also return status which can be
	// // used to break out of loop
	// for {
	// 	msg, open := <-msgChannel

	// 	if !open{
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	//syntactic sugar
	for msg := range msgChannel {
		fmt.Println(msg)
	}
}

//input a channel of type string
func count(thing string, msgChannel chan string) {
	for i := 1; i <= 5; i++ {
		msgChannel <- thing
		time.Sleep(time.Millisecond * 500)

	}

	//since main would be waiting with its open channel
	// this gets detected in runtime
	// to avoid this runtime error we can close the channel
	// A reciever (main) should not close the channel
	// since it can resut in data loss and error
	close(msgChannel)
}
