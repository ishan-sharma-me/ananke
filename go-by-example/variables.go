package main

import "fmt"

func variables() {
	// String type
	var a = "initial"
	fmt.Println("string:", a)

	// Integer types
	var b, c int = 1, 2
	fmt.Println("int:", b, c)

	var d int8 = 127
	fmt.Println("int8:", d)

	var e int16 = 32767
	fmt.Println("int16:", e)

	var f int32 = 2147483647
	fmt.Println("int32:", f)

	var g int64 = 9223372036854775807
	fmt.Println("int64:", g)

	// Unsigned integer types
	var h uint = 1
	fmt.Println("uint:", h)

	var i uint8 = 255
	fmt.Println("uint8:", i)

	var j uint16 = 65535
	fmt.Println("uint16:", j)

	var k uint32 = 4294967295
	fmt.Println("uint32:", k)

	var l uint64 = 18446744073709551615
	fmt.Println("uint64:", l)

	var m uintptr = 0x12345678
	fmt.Println("uintptr:", m)

	// Byte and rune (aliases)
	var n byte = 255 // alias for uint8
	fmt.Println("byte:", n)

	var o rune = 'a' // alias for int32
	fmt.Println("rune:", o)

	// Boolean type
	var p = true
	fmt.Println("bool:", p)

	// Float types
	var q float32 = 3.14159
	fmt.Println("float32:", q)

	var r float64 = 3.141592653589793
	fmt.Println("float64:", r)

	// Complex types
	var s complex64 = 1 + 2i
	fmt.Println("complex64:", s)

	var t complex128 = 1 + 2i
	fmt.Println("complex128:", t)

	// Short variable declarations
	u := "apple"
	fmt.Println("short declaration string:", u)

	v := 42
	fmt.Println("short declaration int:", v)

	w := false
	fmt.Println("short declaration bool:", w)

	some := "thing"
	fmt.Println("short declaration another string:", some)

	// Zero values (uninitialized variables)
	var x int
	var y string
	var z bool
	var aa float64
	var bb complex128
	fmt.Println("zero values - int:", x, "string:", y, "bool:", z, "float64:", aa, "complex128:", bb)
}
