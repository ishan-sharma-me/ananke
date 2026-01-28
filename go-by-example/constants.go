package main

import (
	"fmt"
	"math"
)

func constants() {
	const s string = "constant string"
	fmt.Println(s)

	const n = 10000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
