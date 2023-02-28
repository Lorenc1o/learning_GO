package main

import (
	"fmt"
)

func main() {
	// An if statement can have an optional initialization statement
	// before the condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}
	// The initialization statement is executed before the condition
	// and its variables are only available inside the if statement
	// fmt.Println(num) // undefined: num

	// An if statement can have an optional else statement
	if false {
		fmt.Println("This will not print")
	} else {
		fmt.Println("This will print")
	}

	// An if statement can have an optional else if statement
	if false {
		fmt.Println("This will not print")
	} else if true {
		fmt.Println("This will print")
	}

	// A switch statement is a shorter way to write a sequence of if - else statements
	// It runs the first case whose value is equal to the condition expression
	// If there is no match, it runs the default case
	// It doesn't need a break statement
	// It can have an optional initialization statement
	switch num := 0; {
	case num < 0:
		fmt.Println(num, "is negative")
	case num > 0:
		fmt.Println(num, "is positive")
	default:
		fmt.Println(num, "is zero")
	}

	// A switch statement can have an optional tag
	// The tag can be a variable, a type or a function
	// The tag is evaluated once and the result is compared to the cases
	// If there is no match, it runs the default case
	switch num := 2; num {
	case 0:
		fmt.Println(num, "is zero")
	case 1, 2, 3:
		fmt.Println(num, "is one, two or three")
	default:
		fmt.Println(num, "is not zero or one")
	}

	// If we want fallthrough, we can use the fallthrough keyword
	// It forces the next case to be executed
	// It can be used only in switch statements
	switch num := 2; num {
	case 0:
		fmt.Println(num, "is zero")
	case 1:
		fmt.Println(num, "is one")
		fallthrough
	case 2:
		fmt.Println(num, "is two")
		fallthrough
	case 3:
		fmt.Println(num, "is three")
	default:
		fmt.Println(num, "is not zero or one")
	}

	// We can use the type of a variable in a switch statement
	// It can be used only in switch statements
	var x interface{} = 7
	switch i := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is an integer")
		fmt.Println("x is", i)
	case float64:
		fmt.Println("x is a float64")
	case func(int) float64:
		fmt.Println("x is a function")
	case bool, string:
		fmt.Println("x is a boolean or a string")
	default:
		fmt.Println("I don't know what x is")
	}

}
