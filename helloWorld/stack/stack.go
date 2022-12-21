package stack

import (
	"fmt"
	"strconv"
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
