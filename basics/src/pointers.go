package main

import "fmt"

func main() {
	a := 42 // 42 is the answer to life, the universe, and everything

	fmt.Println(a)
	fmt.Println(&a) // & gives us the memory address of the value this variable is pointing at

	var b *int = &a // *int is a pointer to an int
	fmt.Println(b)
	fmt.Println(*b) // * gives us the value this memory address is pointing at

	*b = 30 // change the value at the memory address b is referencing
	fmt.Println(a)

	a = 50 // change the value at the memory address a is referencing
	fmt.Println(*b)

	// Pointer arithmetic is not allowed in Go
	// b++
	// fmt.Println(*b)

	// The package unsafe can be used to do pointer arithmetic, but it is not recommended

	// We can use the new function to create a pointer to a value
	c := new(int)
	*c = 5
	fmt.Println(c, *c)

	// We can also create a pointer to a struct
	type person struct {
		name string
		age  int
	}
	d := new(person)
	d.name = "John" // dereference is not required (*d).name is automatically dereferenced by the compiler
	d.age = 30
	fmt.Println(d, *d)
	fmt.Println(d.name, d.age)

	// The zero value of a pointer is nil
	var e *int
	fmt.Println(e)
}
