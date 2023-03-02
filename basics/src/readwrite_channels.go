package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(2)
	/*
		go func() {
			i := <-ch
			fmt.Println(i)
			ch <- 27
			wg.Done()
		}()
		go func() {
			ch <- 42
			fmt.Println(<-ch)
			wg.Done()
		}()
	*/                       // It is not a very good idea to have a goroutine sending and receiving values from a channel.
	go func(ch <-chan int) { // A read-only channel
		i := <-ch
		fmt.Println(i)
		//ch <- 27 // This does not work because the channel is read-only.
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // A write-only channel
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()

}
