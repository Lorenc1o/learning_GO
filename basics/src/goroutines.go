package main

import (
	"fmt"
	"time"
)

func main() {
	//go sayHi() // This is a goroutine: a lightweight thread of execution
	// We can use the keyword go to start a goroutine, i.e. it is a function that runs concurrently with the calling one
	//time.Sleep(100 * time.Millisecond) // We can use the Sleep function to wait for a certain amount of time

	var msg = "Hi"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Bye" // We can change the value of the variable and it will affect the goroutine: Race condition
	time.Sleep(100 * time.Millisecond)

	// To avoid this, we can pass the variable as an argument to the goroutine
	msg = "Hi"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Bye"
	time.Sleep(100 * time.Millisecond)
}

func sayHi() {
	fmt.Println("Hi")
}
