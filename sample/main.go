package main

import (
	"fmt"
	"time"

	"github.com/ashuthe1/kuki-memcache/cache"
	"github.com/ashuthe1/kuki-memcache/eviction"
	"github.com/ashuthe1/kuki-memcache/persistence"
)

func main() {
	// Initialize cache
	evictionPolicy := eviction.NewLRU()
	myCache := cache.NewCache(5*time.Minute, 100, evictionPolicy, nil)

	// Add some items to the cache
	myCache.Set("key1", "value1")
	myCache.Set("key2", "value2")

	// Save cache to file
	persistenceManager := persistence.NewFilePersistence("cache_data.json")
	if err := persistenceManager.SaveToFile(myCache.Items()); err != nil {
		fmt.Println("Error saving cache to file:", err)
		return
	}

	fmt.Println("Cache saved to file")

	// Simulate loading cache from file
	loadedItems, err := persistenceManager.LoadFromFile()
	if err != nil {
		fmt.Println("Error loading cache from file:", err)
		return
	}

	// Initialize a new cache with loaded items
	loadedCache := cache.NewCache(5*time.Minute, 100, evictionPolicy, nil)
	loadedCache.SetItems(loadedItems)

	// Retrieve items from loaded cache
	val1, found1 := loadedCache.Get("key1")
	val2, found2 := loadedCache.Get("key2")

	fmt.Println("Loaded cache from file")
	fmt.Printf("key1: %v, found: %v\n", val1, found1)
	fmt.Printf("key2: %v, found: %v\n", val2, found2)
}
