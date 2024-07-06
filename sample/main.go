package main

import (
	"fmt"
	"time"

	"github.com/ashuthe1/kuki-memcache/cache"
	"github.com/ashuthe1/kuki-memcache/eviction"
)

func main() {
	c := cache.NewCache(2*time.Second, 2, eviction.NewLRU(), nil)
	c.Set("key", "Haha")

	// time.Sleep(3 * time.Second)

	val, isExists := c.Get("key")
	if !isExists {
		fmt.Println("Key not found")
		return
	}
	fmt.Println(val)
}
