package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100) // We set the number of CPUs that can be executing simultaneously
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func sayHello() {
	fmt.Println("Hello ", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// Advices:
// - Don't create goroutines in libraries. Let the caller decide how to use them.
// - When creating a goroutine, know how it will end.
// - Use go run -race to analyze race conditions.
