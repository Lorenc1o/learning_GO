package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Instead of using the Sleep function, we can use a WaitGroup to wait for a goroutine to finish
func main() {
	var msg = "Hi"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Bye"
	wg.Wait()
}
