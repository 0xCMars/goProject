package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

type Mercedes struct {
	Car
}

func (this *Car) numberOfWheels() int {
	return this.wheelCount
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

func (this *Mercedes) sayHiToMerkel() {
	fmt.Println("HI Merkel")
}

func main() {
	m := Mercedes{Car{4, nil}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}