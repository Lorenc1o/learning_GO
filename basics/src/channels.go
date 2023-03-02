package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine and receive those values into another goroutine.
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	ch := make(chan int) // We can use the channel to send and receive values of type int.
	wg.Add(2)
	go func() {
		i := <-ch // The <-ch syntax receives a value from the channel.
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // The <-ch syntax sends a value into the channel.
		wg.Done()
	}()
	wg.Wait()

}
