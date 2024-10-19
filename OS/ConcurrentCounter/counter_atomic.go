package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter struct with an atomic integer
type Counter struct {
	count int64
}

// NewCounter creates a new Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Increment increases the count by 1 atomically
func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

// Get returns the current count atomically
func (c *Counter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	var wg sync.WaitGroup

	for {
		var numGoroutines int
		fmt.Println("Enter the number of goroutines to run (0 to exit):")
		fmt.Scan(&numGoroutines)

		if numGoroutines == 0 {
			fmt.Println("Exiting...")
			break
		}

		counter := NewCounter()

		// Start the specified number of goroutines to increment the counter
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < 1000; j++ {
					counter.Increment()
				}
			}()
		}

		wg.Wait() // Wait for all goroutines to finish
		fmt.Println("Final count:", counter.Get())
	}
}
