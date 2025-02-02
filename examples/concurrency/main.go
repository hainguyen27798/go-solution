package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ready bool
	mu    sync.Mutex
	cond  = sync.NewCond(&mu)
)

func waiter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	for !ready {
		cond.Wait() // Wait until notified
	}
	fmt.Printf("Waiter %d proceeding after condition met\n", id)
	mu.Unlock()
}

func announcer() {
	time.Sleep(2 * time.Second) // Simulate work
	mu.Lock()
	ready = true
	mu.Unlock()
	cond.Broadcast() // Notify all waiting goroutines
}

func main() {
	var wg sync.WaitGroup

	// Start several waiters
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go waiter(i, &wg)
	}

	// Start the announcer
	go announcer()

	wg.Wait()
}
