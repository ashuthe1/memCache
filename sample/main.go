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

type Student struct {
	Name string
	age  int
}

func main() {
	// Example: Create a cache with 2 seconds TTL, size of 2 items, and custom eviction policy
	c := cache.NewCache(2*time.Second, 2, NewCustom(), nil)

	totalRegisteredUsers := 100
	c.Set("totalRegisteredUsers", totalRegisteredUsers)

	value, found := c.Get("totalRegisteredUsers")
	if found {
		fmt.Println(value)
	}

	student1 := Student{Name: "Ashutosh", age: 22}
	arr := []int{1, 2, 3, 4, 5}
	c.Set("studentId:1", student1)

	c.Set("arr", arr)

	// Get the value from the cache
	value, found = c.Get("studentId:1")
	if found {
		student := value.(Student)
		fmt.Printf("%+v\n", student)
	}

	value, found = c.Get("arr")
	if found {
		fetchedArr := value.([]int)
		for _, v := range fetchedArr {
			print(v, " ")
		}
		println()
	}
}
