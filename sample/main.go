package main

import (
	"fmt"
	"time"

	"github.com/ashuthe1/kuki-memcache/cache"
	"github.com/ashuthe1/kuki-memcache/eviction"
)

func main() {
	c := cache.NewCache(20*time.Minute, 10, eviction.NewLRU(), nil)

	// Batch set key-value pairs
	items := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": []int{1, 2, 3},
	}
	c.BatchSet(items)

	// Batch get values
	keys := []string{"key1", "key2", "key3"}
	values := c.BatchGet(keys)
	for k, v := range values {
		fmt.Printf("%s: %v\n", k, v)
	}
}
