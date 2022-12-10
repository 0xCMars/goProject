package main

import "fmt"

func main() {
	// 1:
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}
	// 2:
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}

	for i := 0; i < 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	str := "G"
	for i := 1; i < 25; i++ {
		println(str)
		str += "G"
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("the complement of %b is %b\n", i, ^i)
	}

	for i := 1; i < 101; i++ {
		switch {
		case i % 15 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
	
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	LABEL1:	
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}