package main

import (
	"fmt"
	"strconv"
)

// Variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
// The var statement declares a list of variables; as in function argument lists, the type is last.
// Go will infer the type of initialized variables.
// Variables can be declared at package or function level.
var myvar bool = true

// Multiple variables can be declared at once.
var (
	myvar1 bool
	myvar2 string
)

// Lowercase variables are only visible within the package.
var myvar3 bool = true

// Uppercase variables are visible outside the package.
var Myvar4 bool = true

func main() {
	fmt.Println(myvar)
	// The var statement declares a list of variables; as in function argument lists, the type is last.
	var a string = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)
	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)
	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println(f)
	// If a variable is declared to be of a certain type, then you can only assign a value of that type to the variable.
	//f = 2 // Error: cannot use 2 (untyped int constant) as string value in assignment

	// Constants are declared like variables, but with the const keyword.
	const g string = "constant"
	fmt.Println(g)
	// Constants can be character, string, boolean, or numeric values.
	const h = 5000
	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	const i = 30000 / h
	fmt.Println(i)
	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(int64(i))

	// Conversion between types is done using the name of the type to convert to in parentheses.
	var j float64 = float64(i)
	fmt.Println(j)

	// Variables can be used to format strings.
	fmt.Printf("Type: %T Value: %v\n", f, f)

	// Shadowing is allowed.
	var myvar string = "shadowed"
	fmt.Println(myvar)

	// To convert a value to a string, use the built-in function strconv.Itoa.
	var k int = 42
	var l string = strconv.Itoa(k)
	fmt.Println(l)
	// If you try direct convertion:
	l = string(k)
	fmt.Println(l)
}
