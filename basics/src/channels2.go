package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	/*go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()*/ // This is a deadlock. The goroutine only receives a value, but we need 5 values to be sent.
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}
