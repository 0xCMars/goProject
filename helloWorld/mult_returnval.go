package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	sum, mul, div := calculate(10, 2)
	fmt.Println("10 + 2 is ", sum)
	fmt.Println("10 * 2 is ", mul)
	fmt.Println("10 / 2 is ", div)

	fmt.Print("First example with -1: ")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	//you could also write it like this
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(MySqrt(5))

}

func MySqrt(number float64) (result float64, err error) {
	if number < 0 {
		result = float64(math.NaN())
		err = errors.New("number < 0")
	} else {
		result = math.Sqrt(number)
	}
	return 
}

func calculate(a, b int) (sum, mul, div int) {
	sum = a + b
	mul = a * b
	div = a / b
	return 
}