package main

import (
	"fmt"
	"sort"
)

var (
	map1 = map[int]string{
		1:"Monday", 
		2:"Tuesday",
		3:"WednesDay",
		4:"Thursday",
		5:"Friday",
		6:"Saturday",
		7:"Sunday",
	}
	drink = map[string]string{
		"pesi":"百事",
		"coca":"可乐",
	}
)

func main() {
	
	fmt.Println(map1)
	flagHolliday := false

	for k, v := range map1 {
		if v == "thursday" || v == "holliday" {
			fmt.Println(v, " is the ", k, "th day in the week")
			if v == "holliday" {
				flagHolliday = true
			}
		}
	}

	if !flagHolliday {
		fmt.Println("holliday is not a day!")
	}

	fmt.Println(drink)
	for k, v := range drink {
		fmt.Printf("Drink English name is %s, chinese name is: %s\n", k, v)
	}

	keys := make([]string, len(drink))
	i := 0
	for k, _ := range drink {
		keys[i] = k
		i++
	}
	fmt.Println("Now the sorted output:")
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Printf("%s, %s\n", v, drink[v])
	}
}