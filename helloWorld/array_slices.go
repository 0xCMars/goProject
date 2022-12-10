package main
import (
	"fmt"
	"bytes"
)

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
	sl := make([]byte, 5, 10)
	for i := 0; i < 5; i++ {
		sl[i] = 'w'
	}
	s1 := []byte{'v','s','s','s','s','v'}
	sl = Append(sl, s1)
	fmt.Printf("slice = %c \n", sl)

	buffer := bytes.NewBuffer(s1)
	n := 3
	sl1, sl2 := buffer.String()[:n], buffer.String()[n:]
	fmt.Printf("sl1 = %s \n", sl1)
	fmt.Printf("sl2 = %s \n", sl2)


}

func Append(slice []byte, data []byte) []byte {
	lengthNewSlice  := len(slice) + len(data)
	capNewSlice := cap(slice)
    if lengthNewSlice > cap(slice) {
        capNewSlice = lengthNewSlice
    }
    newSlice := make([]byte, lengthNewSlice, capNewSlice)
    for sliceKey, sliceItem := range slice {
        newSlice[sliceKey] = sliceItem
    }
    for dataKey, item := range data {
        newSlice[dataKey + len(slice)] = item
    }
    return newSlice
}