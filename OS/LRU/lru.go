package main

import (
	"container/list"
	"fmt"
)

// LRUCache represents the LRU cache structure
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// Entry holds the key-value pairs stored in the cache
type Entry struct {
	key   int
	value int
}

// NewLRUCache creates a new LRU cache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves the value of a key from the cache
func (lru *LRUCache) Get(key int) (int, bool) {
	if element, found := lru.cache[key]; found {
		// Move the accessed element to the front (most recent)
		lru.list.MoveToFront(element)
		entry := element.Value.(*Entry)
		fmt.Printf("Accessed key %d: %d\n", key, entry.value)
		return entry.value, true
	}
	fmt.Printf("Key %d not found in cache\n", key)
	return -1, false // Return -1 if key not found
}

// Put inserts a key-value pair into the cache
func (lru *LRUCache) Put(key, value int) {
	// Check if key already exists
	if element, found := lru.cache[key]; found {
		// Update the value and move to the front
		lru.list.MoveToFront(element)
		element.Value.(*Entry).value = value
		fmt.Printf("Updated key %d with value %d\n", key, value)
	} else {
		// Insert new entry
		if lru.list.Len() >= lru.capacity {
			// Evict the least recently used (from the back of the list)
			evict := lru.list.Back()
			if evict != nil {
				evictedEntry := evict.Value.(*Entry)
				delete(lru.cache, evictedEntry.key)
				lru.list.Remove(evict)
				fmt.Printf("Evicted key %d\n", evictedEntry.key)
			}
		}
		// Add the new entry to the front of the list
		entry := &Entry{key: key, value: value}
		element := lru.list.PushFront(entry)
		lru.cache[key] = element
		fmt.Printf("Added key %d with value %d\n", key, value)
	}
}

// Display prints the current state of the cache
func (lru *LRUCache) Display() {
	fmt.Print("Cache state: ")
	for element := lru.list.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*Entry)
		fmt.Printf("%d:%d ", entry.key, entry.value)
	}
	fmt.Println()
}

func main() {
	var capacity int
	fmt.Println("Enter the capacity of the LRU cache:")
	fmt.Scan(&capacity)

	cache := NewLRUCache(capacity)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1: Get value by key")
		fmt.Println("2: Put key-value pair")
		fmt.Println("0: Exit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var key int
			fmt.Print("Enter the key to access: ")
			fmt.Scan(&key)
			cache.Get(key)
			cache.Display()

		case 2:
			var key, value int
			fmt.Print("Enter the key: ")
			fmt.Scan(&key)
			fmt.Print("Enter the value: ")
			fmt.Scan(&value)
			cache.Put(key, value)
			cache.Display()

		case 0:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
