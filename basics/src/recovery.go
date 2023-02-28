package main

import (
	"fmt"
)

func main() {
	// The defer statement executes before the panic statement, so the deferred function is executed before the program crashes
	// This allows to recover from the panic
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	panic("a problem")
	fmt.Println("end")
}
