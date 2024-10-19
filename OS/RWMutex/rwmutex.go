package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter struct with an RWMutex
type Counter struct {
	mu    sync.RWMutex
	value int
}

// Increment increases the counter value by 1
func (c *Counter) Increment(id int) {
	c.mu.Lock()         // Lock for writing
	defer c.mu.Unlock() // Unlock after the write
	c.value++
	fmt.Printf("Writer %d incremented the counter to: %d\n", id, c.value)
}

// Value returns the current counter value
func (c *Counter) Value(id int) int {
	c.mu.RLock()         // Lock for reading
	defer c.mu.RUnlock() // Unlock after reading
	fmt.Printf("Reader %d is reading the counter: %d\n", id, c.value)
	time.Sleep(500 * time.Millisecond) // Simulate a longer read delay
	fmt.Printf("Reader %d finished reading\n", id)
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	readerCount := 0
	writerCount := 0

	for {
		fmt.Println("\nEnter 1 to spawn multiple readers, 2 to increment, 0 to exit:")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1: // Spawn multiple readers
			for i := 0; i < 5; i++ { // Spawn 5 concurrent readers
				readerCount++
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					counter.Value(id)
				}(readerCount)
			}

		case 2: // Increment operation
			writerCount++
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				counter.Increment(id)
			}(writerCount)

		case 0: // Exit
			wg.Wait() // Wait for all ongoing operations to finish
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please enter 1, 2, or 0.")
		}
	}
}
