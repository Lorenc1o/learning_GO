package main

import (
	"fmt"
)

// Enumerated constants
const (
	// iota is a predeclared identifier representing the untyped integer ordinal number
	// 	of the current const specification in a (usually parenthesized) const declaration.
	// It is zero-indexed.
	// The first time the identifier is used in a const declaration, it represents the integer 0.
	// Each subsequent use of iota in the same const declaration increments it by one.
	// A const declaration can include multiple iota expressions, one for each constant that needs to be created.
	// The value of each constant is the value of the iota expression at the time the constant is declared.
	a = iota // 0
	b        // 1
	c        // 2
)

// iota is scoped to the constant block.
const (
	d = iota // 0
	e        // 1
	f        // 2
)

// If we don't want the first value to be 0, we can use the following syntax.
const (
	g = iota + 1 // 1
	h            // 2
	i            // 3
)

// If we don't want to store the value of iota, we can use the underscore character.
const (
	j = iota // 0
	_        // 1
	k        // 2
)

// We can also use iota to create a bit mask.
const (
	// 1 << iota is a binary number that is 1 shifted left by iota.
	// 1 << 0 == 00000001
	// 1 << 1 == 00000010
	// 1 << 2 == 00000100
	// and so on...
	l = 1 << iota // 1
	m             // 2
	n             // 4
	o             // 8
)

func main() {
	// Constant are named in the same way as variables.
	const myconst bool = true // internal scoped
	const MyConst bool = true // external scoped
	// Constants can be character, string, boolean, or numeric values.
	// They need to be initialized at declaration and at compile time.
	var role byte = l | m | n
	fmt.Printf("Role: %b\n", role)
}
