package main

import "fmt"

func values() {
	// Basic string concatenation
	fmt.Println("go" + "lang")

	// Basic arithmetic
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Basic boolean logic
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Complex arithmetic operations
	fmt.Println("2*3+4 =", 2*3+4)
	fmt.Println("10-5*2 =", 10-5*2)
	fmt.Println("17%5 =", 17%5)
	fmt.Println("2^3 =", 1<<3) // Using bitwise shift instead of ^ which is XOR

	// Floating point precision issues
	fmt.Println("0.1+0.2 =", 0.1+0.2)
	fmt.Println("0.1+0.2 == 0.3:", 0.1+0.2 == 0.3)

	// Large numbers and overflow considerations
	var bigInt int64 = 9223372036854775807
	fmt.Println("Max int64:", bigInt)
	fmt.Println("Max int64 + 1:", bigInt+1) // This will overflow

	// Negative numbers and operations
	fmt.Println("-5 + 3 =", -5+3)
	fmt.Println("5 * -3 =", 5*-3)
	fmt.Println("-7 / 3 =", -7/3) // Integer division towards zero

	// Bitwise operations
	fmt.Println("5 & 3 =", 5&3)     // AND
	fmt.Println("5 | 3 =", 5|3)     // OR
	fmt.Println("5 ^ 3 =", 5^3)     // XOR
	fmt.Println("5 &^ 3 =", 5&^3)   // AND NOT
	fmt.Println("1 << 3 =", 1<<3)   // Left shift
	fmt.Println("16 >> 2 =", 16>>2) // Right shift

	// Complex boolean logic
	fmt.Println("true && false || true =", true && false || true)
	fmt.Println("false || true && false =", false || true && false)
	fmt.Println("!(true && false) =", !(true && false))
	fmt.Println("(true || false) && (false || true) =", (true || false) && (false || true))

	// Comparison operations
	fmt.Println("5 > 3 =", 5 > 3)
	fmt.Println("5 >= 5 =", 5 >= 5)
	fmt.Println("3 < 5 =", 3 < 5)
	fmt.Println("5 <= 5 =", 5 <= 5)
	fmt.Println("5 == 5 =", 5 == 5)
	fmt.Println("5 != 3 =", 5 != 3)

	// String operations
	fmt.Println(`"hello" + " world" =`, "hello"+" world")
	fmt.Println(`len("golang") =`, len("golang"))

	// Complex number operations
	c1 := 1 + 2i
	c2 := 3 + 4i
	fmt.Println("Complex addition:", c1+c2)
	fmt.Println("Complex multiplication:", c1*c2)
	fmt.Println("Complex division:", c1/c2)
	fmt.Println("Real part of c1:", real(c1))
	fmt.Println("Imaginary part of c1:", imag(c1))

	// Type conversion examples
	var x int = 42
	var y float64 = float64(x)
	fmt.Println("int to float64:", x, "->", y)

	var a float64 = 3.14159
	var b int = int(a)
	fmt.Println("float64 to int (truncation):", a, "->", b)

	// Mixed type operations (requires explicit conversion)
	var i int = 5
	var f float64 = 2.5
	fmt.Println("Mixed types with conversion:", float64(i)*f)

	// Precedence confusion examples
	fmt.Println("2+3*4 =", 2+3*4)
	fmt.Println("(2+3)*4 =", (2+3)*4)
	fmt.Println("true && false || true && false =", true && false || true && false)
	fmt.Println("true && (false || true) && false =", true && (false || true) && false)

	// Edge cases with floating point precision
	fmt.Println("0.1 + 0.2 =", 0.1+0.2)
	fmt.Println("Floating point precision check:", 0.1+0.2 == 0.3)

	// String formatting that might be confusing
	fmt.Printf("Formatted: %d, %f, %t, %s\n", 42, 3.14159, true, "golang")
	fmt.Printf("Binary: %b, Hex: %x, Octal: %o\n", 42, 42, 42)
}
