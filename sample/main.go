package main

import (
	"fmt"
	"time"

	"github.com/ashuthe1/kuki-memcache/cache"
	"github.com/ashuthe1/kuki-memcache/eviction"
)

func main() {
	// Create a cache with default TTL of 5 minutes, size of 10 items, and LRU eviction policy
	c := cache.NewCache(5*time.Minute, 10, eviction.NewLRU(), nil)

	// Set a key-value pair with a custom TTL of 10 seconds
	c.Set("keyWithTTL", "value1", 10*time.Second)

	// Set a key-value pair with the default TTL
	c.Set("keyWithDefaultTTL", "value2")

	// Retrieve and print the values before expiration
	val, found := c.Get("keyWithTTL")
	if found {
		fmt.Println("keyWithTTL:", val)
	} else {
		fmt.Println("keyWithTTL: not found or expired")
	}

	val, found = c.Get("keyWithDefaultTTL")
	if found {
		fmt.Println("keyWithDefaultTTL:", val)
	} else {
		fmt.Println("keyWithDefaultTTL: not found or expired")
	}

	// Wait for 11 seconds to let the custom TTL item expire
	time.Sleep(11 * time.Second)

	// Retrieve and print the values after expiration
	val, found = c.Get("keyWithTTL")
	if found {
		fmt.Println("keyWithTTL:", val)
	} else {
		fmt.Println("keyWithTTL: not found or expired")
	}

	val, found = c.Get("keyWithDefaultTTL")
	if found {
		fmt.Println("keyWithDefaultTTL:", val)
	} else {
		fmt.Println("keyWithDefaultTTL: not found or expired")
	}
}
