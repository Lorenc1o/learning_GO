package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 50) // A buffered channel. It can hold 50 values.
	wg.Add(2)
	go func(ch <-chan int) { // A read-only channel
		for i := range ch { // We can use a for range loop to receive values from a channel until it is closed.
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // A write-only channel
		ch <- 42
		ch <- 27
		close(ch) // We need to close the channel to indicate that no more values will be sent.
		// ch <- 33 // This does not work because the channel is closed.
		wg.Done()
	}(ch)
	wg.Wait()
}
