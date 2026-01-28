package main

import "fmt"

func forExample() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	println()

	for _, v := range []string{"a", "b", "c"} {
		fmt.Print(v)
	}

	println()

	for i := range 20 {
		fmt.Print(i)
	}

	println()

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Print(i)
	}

	println()

	for i := range 6 {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}
}
