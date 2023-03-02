package main

import (
	"fmt"
	"sync"
)

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.
// A default case in a select is run if no other case is ready.
var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	quit := make(chan struct{}) // A struct{} is a zero-byte struct. It is used to signal a channel to close.
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			select {
			case i := <-ch:
				fmt.Println(i)
			case <-quit:
				wg.Done()
				return
				// default: // If we uncomment this, the program will not wait until something is sent into the channel.
			}
		}
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		quit <- struct{}{} // We send a value into the channel to signal that we are done.
		wg.Done()
	}(ch)
	wg.Wait()
}
