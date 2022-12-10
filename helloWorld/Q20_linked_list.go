package main

import (
	"fmt"
	"container/list"
	"unsafe"
)

func main() {
	a := list.New()
	a.PushBack(101)
	a.PushBack(102)
	a.PushBack(103)
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Printf("Numnber is: %d\n", i.Value)
	}

	var i int = 19
	size := unsafe.Sizeof(i)
	fmt.Println("The size of an int is: ", size)
}