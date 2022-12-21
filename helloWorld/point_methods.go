package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Point3D struct {
	X, Y, Z float64
}

type Polar struct {
	R, T float64
}

type Magnitude interface {
	Abs() float64
}

var m Magnitude

func (p *Point) Abs() float64 {
	return  math.Sqrt(float64(p.X * p.X + p.Y * p.Y))
}

func (p *Point3D) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

func (p Polar) Abs() float64 { return p.R }

func (p *Point) Scale(sclar float64) {
	p.X = p.X * sclar
	p.Y = p.Y * sclar
	return
} 

func main() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs())

	p2 := &Point{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", p2.Abs())

	p1.Scale(5)
	m = p1
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", m.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f\n", p1.X, p1.Y)

	mag := m.Abs()
	m = &Point3D{3, 4, 5}
	mag += m.Abs()
	m = Polar{2.0, math.Pi / 2}
	mag += m.Abs()
	fmt.Printf("The float64 mag is now: %f", mag)
}