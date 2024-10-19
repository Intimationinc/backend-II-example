package main

import (
	"fmt"
	"sync"
	"time"
)

// Job struct represents a task that needs to be processed
type Job struct {
	id int
}

// Worker function processes a job and simulates work with sleep
func worker(job Job, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire semaphore slot
	semaphore <- struct{}{}

	// Simulate job processing
	fmt.Printf("Processing job ID: %d\n", job.id)
	time.Sleep(1 * time.Second) // Simulate some processing time

	// Release semaphore slot
	<-semaphore

	fmt.Printf("Job ID: %d processed\n", job.id)
}

// Function to process jobs from a queue with a limited number of workers
func processQueue(jobs []Job, maxWorkers int) {
	var wg sync.WaitGroup

	// Create a semaphore with maxWorkers slots
	semaphore := make(chan struct{}, maxWorkers)

	// Start processing each job
	for _, job := range jobs {
		wg.Add(1)
		go worker(job, semaphore, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
}

func main() {
	// Simulate a queue of jobs
	jobs := []Job{
		{1}, {2}, {3}, {4}, {5},
		{6}, {7}, {8}, {9}, {10},
	}

	// Set max number of workers (concurrent job processors)
	maxWorkers := 3

	fmt.Println("Starting to process the job queue...")
	processQueue(jobs, maxWorkers)
	fmt.Println("All jobs have been processed.")
}
