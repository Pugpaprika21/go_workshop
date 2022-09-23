package main

import (
	"fmt"
)

func getArray(m map[string]int) {
	for k, v := range m {
		if m["one"] == 5 {
			fmt.Println("find Key and Value done!", k, v)
		} else {
			fmt.Println("find Key and Value fail!", k, v)
		}
	}
}

func main() {

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	getArray(m)
}
