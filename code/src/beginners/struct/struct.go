// learn using structs
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Hero", age: 32}
	fmt.Println(p.age)
}
