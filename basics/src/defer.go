package main

import (
	"fmt"
)

func main() {
	// A defer statement defers the execution of a function until the surrounding function returns
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
	// Deferred function calls are pushed onto a stack
	// When a function returns, its deferred calls are executed in last-in-first-out order
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("!")

	a := "start"
	defer fmt.Println(a) // The deferred call's arguments are evaluated immediately,
	//  but the function call is not executed until the surrounding function returns
	a = "end"
}
