package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { // recover is a built-in function that regains control of a panicking goroutine
			log.Println("Error:", err)
			panic(err) // If we don't know how to handle the error, we can re-panic
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
