package main

import (
	"fmt"
	"math"
)

const s string = "constants"

func main() {

	fmt.Println(s)
	const x = 1000
	const y = 2000 / x
	fmt.Println(math.Sin(y))
}
