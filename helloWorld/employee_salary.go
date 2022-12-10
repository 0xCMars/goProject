package main

import (
	"fmt"
)

type employee struct {
	salary float64
}

func (emp *employee) giveRaise(per float64) {
	emp.salary = (1 + per) * emp.salary
}

func main() {
	/* create an employee instance */
	var e = new(employee)
	e.salary = 100000
	/* call our method */
	e.giveRaise(0.04)
	fmt.Printf("Employee now makes %f", e.salary)
}
