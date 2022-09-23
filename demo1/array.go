package main

import (
	"fmt"
)

func getArray(param []int, counter int) float32 {
	var result int
	var total float32
	for i := 0; i < counter; i++ {
		result += param[i]
	}
	total = float32(result / counter)
	return total
}

func main() {

	array := []int{20, 30, 40, 50}
	slice := make([]int, 10, 20)
	append := append(slice, 20, 55)
	fmt.Println(getArray(array, 4))
	fmt.Println(len(array))
	fmt.Println(cap(array))
	fmt.Println(slice)
	fmt.Println(append)

	/* var way1 [4]string
	way1[0] = "Hello"
	way1[1] = "World"

	way2 := [4]int{2, 1, 3, 4}
	way3 := []int{10, 20, 30, 40}

	fmt.Println(way1)
	fmt.Println(way2)
	fmt.Println(way3) */
}
