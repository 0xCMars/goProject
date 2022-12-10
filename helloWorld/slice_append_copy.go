package main
import "fmt"

func main() {
	s := []int{1,2,3,4,5,6}
	factor := 5

	newSlice := make([]int, len(s) * factor)
	n := copy(newSlice, s)
	fmt.Println(newSlice)
	fmt.Printf("Copied %d elements\n", n)

	s = Filter(s, event)
	fmt.Println(s)

	a := []string{"M", "N", "O", "P", "Q", "R"}
	b := []string{"A", "B", "C"}
	res := InsertStringSlice(a, b, 0) // at the front
	fmt.Println(res)                   // [A B C M N O P Q R]
	res = InsertStringSlice(a, b, 3)  // [M N O A B C P Q R]
	fmt.Println(res)

	result := RemoveStringSlice(2,4, a)
	fmt.Println(result)
	fmt.Println(a)
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, val := range s {
		if fn(val) {
			p = append(p,val)
		}
	}
	return p
}

func event(n int) bool {
	if n%2 == 0{
		return true
	} else {
		return false
	}
}

func InsertStringSlice(to, from []string, pos int) []string {
	var p []string

	p = append(p, to[:pos]...)
	p = append(p, from...)
	p = append(p, to[pos:]...)
	return p
}

func RemoveStringSlice(start, end int, slice []string) []string {
	var p []string = make([]string, len(slice) - (end - start))

	at := copy(p, slice[:start])
	at += copy(p[at:], slice[end:])

	return p
}