package main

import (
	"fmt"
)

func main() {
	// A panic is a built-in function that stops the ordinary flow of control and begins panicking
	// When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller
	// To the caller, F then behaves like a call to panic
	// The process continues up the stack until all functions in the goroutine have returned, at which point the program crashes
	// Panics can be initiated by invoking panic directly
	// They can also be caused by runtime errors, such as out-of-bounds array accesses
	fmt.Println("start")
	panic("a problem")
	fmt.Println("end")

	// A deferred function's arguments are evaluated when the defer statement is evaluated
	// fmt.Println("deferred") // This line will not be executed
	// A panic can be recovered from
	// A call to recover stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic
	// If recover is called outside the deferred function it will not stop a panicking sequence
	// In this example, createFile will panic, but because we defer the call to closeFile, closeFile will be called before the program exits
	// f := createFile("/tmp/defer.txt")
	// defer closeFile(f)
	// writeFile(f)
}
