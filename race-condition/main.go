package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var counter int

func main() {
	for range 1000 {
		wg.Go(increment)
		// go increment()
	}

	wg.Wait()

	fmt.Println("counter", counter)
}

func increment() {
	defer mu.Unlock()
	mu.Lock()
	counter++
}
