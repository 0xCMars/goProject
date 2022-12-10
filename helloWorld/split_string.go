package main
import (
	"fmt"
)

func main() {
	str := "Google"
	for i := 0; i <= len(str); i++ {
		a, b := SplitString(str, i)
		fmt.Printf("The string %s split at position %d is: %s / %s\n", str, i, a, b)
	}

	result := SplitString2(str)
	fmt.Printf("Result is: %s\n", result)

	Q29 := "abaaacdefg"
	result2 := Q29Uniq(Q29)
	fmt.Printf("Result is: %s\n", string(result2))

	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", sla)
	// sla is passed via call by value, but since sla is a reference type
	// the underlying slice is array is changed (sorted)
	bubbleSort(sla)
	fmt.Println("after sort:  ", sla)

	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i * 10
	}

	println()
	// shorter:
	fmt.Printf("%v\n", MapFunc(mf, list))
}

func SplitString(str string, i int) (first, second string) {
	first = str[:i]
	second = str[i:]
	return
}

func SplitString2(str string) (result string) {
	result = str[len(str)/2:] + str[:len(str)/2]
	return
}

func StringReverse(str string) (result string) {
	b := []rune(str)
	n, mid := len(b), len(b)/2 
	for i := 0; i < mid; i++ {
		b[i], b[n-i-1] = b[n-1-i], b[i]
	}
	return string(b)
}

func Q29Uniq(str string) (result []rune) {
	runes := []rune(str)
	n := len(runes)
	for i := 1; i < n; i++ {
		if runes[i] != runes[i-1] {
			result = append(result, runes[i])
		}
	}
	return
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < n - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func MapFunc(fn func(int) int, arr []int) []int {
	result := make([]int, len(arr))
	for i, val := range arr {
		result[i] = fn(val)
	}
	return result
}

func Multi10(num int) int {
	return num * 10;
}