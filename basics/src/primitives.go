package main

import (
	"fmt"
	"strings"
)

func main() {
	// Booleans
	// Boolean values can be either true or false.
	var mybool bool = true
	fmt.Printf("Type: %T Value: %v\n", mybool, mybool)
	// Booleans can be created by using a predicate expression.
	mybool2 := 1 == 1
	fmt.Printf("Type: %T Value: %v\n", mybool2, mybool2)
	// An uninitialized boolean is false.
	var mybool3 bool
	fmt.Printf("Type: %T Value: %v\n", mybool3, mybool3)

	// Integers
	// Integer can be int8, int16, int32 and int64.
	// Unsigned integers can be uint8, uint16, uint32 and uint64.
	// The basic operations (they only work with integers of the same size):
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	// Bitwise operations:
	fmt.Println(a & b)  // Bitwise AND
	fmt.Println(a | b)  // Bitwise OR
	fmt.Println(a ^ b)  // Bitwise XOR
	fmt.Println(a &^ b) // Bit clear (AND NOT)
	// Shifting bits:
	fmt.Println(a << 3) // Shift bits to the left (multiply by 2^3)
	fmt.Println(a >> 3) // Shift bits to the right (divide by 2^3)

	// Floating-point numbers
	// Floating-point numbers can be float32 and float64.
	// The basic operations (they only work with floating-point numbers of the same size):
	c := 10.2
	d := 0.37e2
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)

	// Complex numbers
	// Complex numbers can be complex64 and complex128.
	// The basic operations (they only work with complex numbers of the same size):
	e := 1 + 2i
	f := 3 + 4i
	fmt.Println(e + f)
	fmt.Println(e - f)
	fmt.Println(e * f)
	fmt.Println(e / f)
	// The real and imaginary parts of a complex number can be extracted using the real and imag functions.
	fmt.Println(real(e))
	fmt.Println(imag(e))
	// Complex numbers can be created using the complex function.
	g := complex(1, 2)
	fmt.Println(g)

	// Strings
	// Strings are immutable sequences of bytes.
	// Strings can be created using double quotes or backticks.
	s1 := "Hello World"
	s2 := `Hello World`
	// Strings can be concatenated using the + operator.
	s3 := "Hello" + " " + "World"
	// Strings can be concatenated using the += operator.
	s3 += "!"
	// Strings can be concatenated using the fmt package.
	s4 := fmt.Sprintf("%s %s!", s1, s2)
	// Strings can be concatenated using the strings package.
	s5 := strings.Join([]string{s1, s2, s3}, " ")
	fmt.Println(s4)
	fmt.Println(s5)

	// Rune
	// A rune is an alias for int32 and is equivalent to a Unicode code point.
	// A rune can be created using single quotes.
	r := 'a'
	fmt.Printf("Type: %T Value: %v\n", r, r)
	// A rune can be converted to a string using the string function.
	s := string(r)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	// A string can be converted to a rune using the []rune function.
	r2 := []rune(s)
	fmt.Printf("Type: %T Value: %v\n", r2, r2)

}
