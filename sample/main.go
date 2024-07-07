package main

import (
	"fmt"
	"time"

	"github.com/ashuthe1/kuki-memcache/cache"
	"github.com/ashuthe1/kuki-memcache/eviction"
)

type CustomEviction struct {
	// Define your custom data structure and fields here
}

func NewCustom() eviction.EvictionPolicy {
	return &CustomEviction{
		// Initialize your custom data structure here
	}
}

func (c *CustomEviction) Add(key string) {
	// Implement your custom eviction policy logic for Add operation
}

func (c *CustomEviction) Remove(key string) {
	// Implement your custom eviction policy logic for Remove operation
}

func (c *CustomEviction) Evict() string {
	// Implement your custom eviction policy logic for Evict operation
	return "" // Adjust the return value as per your implementation
}

func main() {
	// Example: Create a cache with 2 seconds TTL, size of 2 items, and custom eviction policy
	c := cache.NewCache(2*time.Second, 2, NewCustom(), nil)

	// Set a key-value pair in the cache
	c.Set("key", "value")

	// Get a value from the cache
	val, isExists := c.Get("key")
	if !isExists {
		fmt.Println("Key not found")
	}

	fmt.Println(val)

	time.Sleep(3 * time.Second)
	val, isExists = c.Get("okkkk")
	if !isExists {
		fmt.Println("Key not found")
	}

	fmt.Println(val)

	fmt.Println(c.Expired())
	fmt.Println(c.Hits())
	fmt.Println(c.Misses())
}
