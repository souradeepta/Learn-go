// Create a calculator
// short hand init declarations
// x:= 5

package main

import "fmt"

func main() {

	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x)
	fmt.Scan(&y)

	var sum int = x + y
	var sub int = x - y

	fmt.Println("Sum, Diff is: ", sum, sub)
}
