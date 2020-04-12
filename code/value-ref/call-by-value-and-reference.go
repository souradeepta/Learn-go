// Showing call by value and call
// by reference
package main

import (
	"fmt"
)

func main() {

	var i int
	fmt.Print("input a integer: ")
	fmt.Scan(&i)

	byValue(i)
	fmt.Println("incrementing by value", i)

	//send pointer with &
	byRef(&i)
	fmt.Println(" incrementing by ref", i)
}

func byValue(x int) {
	x++
}

func byRef(x *int) {
	//deference the pointer
	*x++
}
