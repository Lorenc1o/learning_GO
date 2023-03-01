package main

import "fmt"

// we can define a function with the func keyword
// if it is named with lower case, it is private
// if it is named with upper case, it is public to the package
// we can define the return type after the arguments
func add(x int, y int) int {
	return x + y
}

// we can define multiple arguments of the same type
func add2(x, y int) int {
	return x + y
}

// the parameters are passed by value
// if we want to pass by reference, we can use pointers
// we can return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// We can define a function with a variable number of arguments
// The arguments are passed as a slice
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// We can name the return values
// If we name the return values, we can use the return statement without arguments
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// We can return errors
// The error type is a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, fmt.Errorf("can't work with 42")
	}
	return arg + 3, nil
}

func main() { // main is the entry point of the program
	fmt.Println(add(42, 13)) // we can call the function with the name and the arguments
	fmt.Println(add2(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(split(17))
	fmt.Println(f1(42))
	fmt.Println(f1(43))

	// Anonymous functions
	// We can define a function without a name
	// We can use it as a value
	func() {
		fmt.Println("hello GO!")
	}()
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// We can use functions as variables
	f := func() {
		fmt.Println("hello GO!")
	}
	f()
}
