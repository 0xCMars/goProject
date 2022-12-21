package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	i int
}

type RSimple struct {
	w int
}

func (this *RSimple) Get() int {
	return this.w
}

func (this *RSimple) Set(i int) {
	this.w = i
}

func (this *Simple) Get() int {
	return this.i
}

func (this *Simple) Set(i int) {
	this.i = i
}

// func fI(it Simpler) int {
// 	it.Set(5)
// 	return it.Get()
// }

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple
	fmt.Println(fI(&r))
}
