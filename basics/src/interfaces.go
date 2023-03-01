package main

import "fmt"

func main() {
	// We can define a variable of an interface type
	var w Writer

	// We can assign a value of a concrete type that implements the interface
	w = ConsoleWriter{}

	// We can call the method on the interface
	w.Write([]byte("Hello Go!"))
}

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type Writer interface {
	Write([]byte) (int, error) // Write is a method that takes a byte slice and returns an int and an error
}

// We can define a type that implements the Writer interface
type ConsoleWriter struct{}

// We can define a method on the ConsoleWriter type
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
