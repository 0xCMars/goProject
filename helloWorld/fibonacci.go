package main

import "fmt"

func main() {
	result := 0
	index := 0
	for i := 0; i <= 10; i++ {
		result, index = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		fmt.Printf("fibonacci(%d) is: %d\n", index, result)
	}
	Ten2One(1)
	x := 0
	// x = factorial(30)
	fmt.Printf("30-th factorial result is: %d\n", x)

	fv := func() { fmt.Println("Hello World.") }
	fv()
	fmt.Printf("fv is of type %T and has value %s\n", fv, fv)
	fv()
	fmt.Printf("fv is of type %T and has value %s\n", fv, fv)

	fi_closure := fibonacci_closure() 
	for i := 0; i <= 9; i++ {
		result = fi_closure()
		fmt.Printf("fibonacci(%d) is: %d\n", i+ 2, result)
	}
}

func fibonacci(n int) (res int, index int) {
	if n <= 1 {
		res = 1 
	} else {
		c1, _ := fibonacci(n-1) 
		c2, _ := fibonacci(n-2)
		res = c1 + c2
	}
	index = n
	return
}

func Ten2One(n int) {
	if n == 10 {
		fmt.Println("%d\n", n)
	} else {
		Ten2One(n+1)
		fmt.Println("%d\n", n)
	}
}

func factorial(n int) (result int){
	if n <= 0 {
		result = 1
		fmt.Printf("%d-th factorial is %d\n", n, result)
	} else {
		result = n * factorial(n-1)
		fmt.Printf("%d-th factorial is %d\n", n, result)
	}
	return
}

func fibonacci_closure() func() int {
	var first int = 1
	var second int = 1
	return func() int {
		first, second = second, first + second
		return second
	}
}