package main

import (
	"fmt"
)

type Rectangle struct {
	length, width int
}

func Area(r *Rectangle) (area int) {
	return r.length * r.width
}

func Perimeter(r *Rectangle) (perimeter int) {
	perimeter = 2 * (r.length + r.width )
	return 
}

func main() {
	r1 := new(Rectangle)
	r1.length = 10
	r1.width = 5 
	fmt.Printf("The area of r1 is: %d\n", Area(r1))
	fmt.Printf("The Perimeter of r1 is: %d\n", Perimeter(r1))

}