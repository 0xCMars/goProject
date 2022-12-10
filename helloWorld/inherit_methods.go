package main

import (
	"fmt"
)

type Base struct {
	id int
}

func (this *Base) Id() int {
	return this.id
}

func (this *Base) SetId(s int) {
	this.id = s
}

type Person struct {
	Base
	FirstName, LastName string
}

type Employee struct {
	Person
	salary int
}

func main() {
	em := new(Employee)
	em.Person = Person{Base{1}, "Mars", "Chen"}
	em.salary = 1000
	fmt.Printf("id is %d\n", em.Id())
}