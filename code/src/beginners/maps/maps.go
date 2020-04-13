// using maps
package main

import (
	"fmt"
)

func main() {
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	delete(vertices, "square")
	fmt.Println(vertices)
}
