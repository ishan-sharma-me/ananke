package main

import (
	"fmt"
	"maps"
)

func mapsExample() {
	fmt.Println("maps example")

	x := make(map[string]int)
	x["k1"] = 10
	x["k2"] = 20
	fmt.Println(x)
	fmt.Println("get k2:", x["k2"])
	fmt.Println("len:", len(x))

	v1 := x["k1"]
	v2 := x["k3"]

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	_, prs := x["k2"]

	fmt.Println("prs:", prs)

	delete(x, "k2")

	fmt.Println(x["k2"])

	clear(x)

	fmt.Println(x)

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
