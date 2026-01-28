package main

import "fmt"

func arraysExample() {

	var a [5]int
	fmt.Println("array:", a)

	a[0] = 42
	fmt.Println(a)

	fmt.Println("length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array literal:", b)

	c := [...]string{"foo", "bar", "baz"}

	fmt.Println("array literal with ... :", c)

	d := [...]int{1: 10, 3: 30}
	fmt.Println("array literal with index specified:", d)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d array:", twoD)

	twoD = [2][3]int{
		{9, 8, 7},
		{6, 6, 6},
	}

	fmt.Println("2d array literal:", twoD)

}
