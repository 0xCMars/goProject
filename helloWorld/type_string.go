package main

import (
	"fmt"
	"strconv"
)

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

type T struct {
    a int
    b float32
    c string
}

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

type Celsius float64

type Day int

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time"}

func (d Day) String() string {
	return dayName[d]
}

func (tz TZ) String() string { // Method on TZ (not ptr)
	if zone, ok := timeZones[tz]; ok {
		return zone
	}
	return ""
}

func (c Celsius) String() string {
	return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

func (t *T) String() string {
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

func main() {
	test := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", test)

	var c Celsius = 18.36
	fmt.Println(c)

	var th Day = 3
	fmt.Printf("The 3rd day is: %s\n", th)
	// If index > 6: panic: runtime error: index out of range
	// but use the enumerated type to work with valid values:
	var day = SU
	fmt.Println(day) // prints Sunday
	fmt.Println(0, MO, 1, TU)

	fmt.Println(EST) // Print* knows about method String() of type TZ
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}