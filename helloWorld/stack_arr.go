package main

import (
	"fmt"
	"strconv"
	"runtime"
)

type Stack struct{
	data [4]int
	top  int
} 

func (s *Stack) Push(i int) {
	if (s.top >= 4) {
		fmt.Println("Stack Full")
		return
	} else {
		s.data[s.top] = i
		s.top += 1
	}
}

func (s *Stack) Pop() (res int) {
	if s.top < 0 {
		return 0
	} else {
		res = s.data[s.top - 1]
		s.top -= 1
		return
	}
}

func (s *Stack) String() string {
	str := ""
	for i := s.top - 1; i>=0; i-- {
		str = "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]" + str
	}
	return str
}

func main() {
	st1 := new(Stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}