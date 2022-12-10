package main

import (
	"fmt"
)

type Anony struct {
	f float32
	int
	string
}

func main() {
	a := new(Anony)
	a.f = 3.2
	a.int = 5
	a.string = "right"

	fmt.Printf("a.f is %f\n", a.f)
	fmt.Printf("a.int is %d\n", a.int)
	fmt.Printf("a.string is %s\n", a.string)
	
}