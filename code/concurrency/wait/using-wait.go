// using go routines for concurrency
package main

import (
	"fmt"
	"sync"
	"time"
)

//says main is redeclared in this block
// since there can only be one main in a package/directory
// fixed: created a new folder for it
func main() {
	//like a counter
	var wg sync.WaitGroup
	//wait for one go routine
	wg.Add(1)

	//wrapping call count in
	// an anonymous function
	// which immediately runs it
	// This is an inline function
	// and so has access to wg
	go func() {
		count("sheep")
		//decrements counter by 1
		wg.Done()
	}()

	// polling wait at the end of main
	// waits until counter is 0
	// waits for go routines to
	// complete
	wg.Wait()

}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
