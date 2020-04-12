package main

import "fmt"

func main() {
	// var a[5]int
	// note that arrays are immutable cannot
	// be modified
	a := [5]int{5, 4, 3, 2, 1}

	// instead you can create mutable slices
	b := []int{1, 2, 3}
	b = append(b, 12)

	fmt.Println(a, b)

}
