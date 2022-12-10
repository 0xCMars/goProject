package main

import (
	"fmt"
)

func f(a [3]int) { 
	a[1] = 10
	fmt.Println(a) 
}
func fp(a *[3]int) { 
	a[1] = 10
	fmt.Println(a) 
}

func main() {
	var ar [3]int
	f(ar) 	// passes a copy of ar
	fmt.Println(ar)
	fp(&ar) // passes a pointer to ar
	fmt.Println(ar)

	arr := new([15]int)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)
}