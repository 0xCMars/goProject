package main
import (
	"unicode/utf8"
	"fmt"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"
	// test := "こん"
	// fmt.Printf("test len is %d\n", len(test))
	// fmt.Printf("The number of characters in string test is %d\n", utf8.RuneCountInString(test))

	// it take three bytes to represengt each char in こん, so the length of str2 is 6 bytes longer the str1's
	// but in the concept of rune, こん mean two unicode, so the rune of str2 is 2 more than str1's
	fmt.Printf("str1 len is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 len is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))


}