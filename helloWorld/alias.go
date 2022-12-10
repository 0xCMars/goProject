package main
import "fmt"

type TZ int
type Rope string
// type Integer int
// func (p Integer) get() int { 
// 	return int(p)
// }

type Integer struct {
	n int
}

func (p Integer) get() int { 
	return p.n
}

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d\n", c) // 输出：c has the value: 7
	// var hello Rope = "hahahaha"
	// // cant compile because alias doesnt have string property
	// fmt.Printf(hello)

	var v *Integer = new(Integer);
	v.n = 10
	fmt.Printf("get result is: %d\n", v.get()) 
}