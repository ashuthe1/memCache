# In-Memory Caching Library

This project implements an in-memory caching library in Golang with support for multiple eviction policies (FIFO, LRU, LIFO) and custom eviction policies.

## Features

- **Eviction Policies**: Supports FIFO, LRU, and LIFO eviction policies.
- **Custom Eviction Policy**: Allows users to define and integrate custom eviction policies.
- **Thread Safety**: Utilizes `sync.Mutex` for concurrent access safety.
- **Statistics**: Tracks cache hits, misses, and expired items.
- **Benchmarking**: Basic benchmarking capabilities for cache operations.

## Setup

### Prerequisites

- Golang installed on your machine

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ashuthe1/kuki-memcache
   cd kuki-memcache
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

### Importing the Library

Import the library into your Go project:

```go
import (
	"time"
	"fmt"

	"github.com/ashuthe1/kuki-memcache/cache"
	"github.com/ashuthe1/kuki-memcache/eviction"
)
```

### Creating and Using the Cache

Initialize a cache instance with your desired TTL (time-to-live), maximum size, and eviction policy:

```go
func main() {
	// Example: Create a cache with 2 seconds TTL, size of 2 items, and LRU eviction policy
	c := cache.NewCache(2*time.Second, 2, eviction.NewLRU(), nil)

	// Set a key-value pair in the cache
	c.Set("key", "value")

	// Get a value from the cache
	val, isExists := c.Get("key")
	if !isExists {
		fmt.Println("Key not found")
		return
	}
	fmt.Println(val)
}
```

### Creating a Custom Eviction Policy

To create your own custom eviction policy, implement the `EvictionPolicy` interface:

```go
type EvictionPolicy interface {
	Add(key string)
	Remove(key string)
	Evict() string
}
```

Here is an example of how to implement and use a custom eviction policy:

```go
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
		return
	}
	fmt.Println(val)
}
```

### Supported Operations

- **Set**: Add a key-value pair to the cache.
- **Get**: Retrieve a value associated with a key from the cache.
- **Delete**: Remove a key-value pair from the cache.
- **Statistics**: Track cache hits, misses, and expired items.

### Running Tests

Run tests to ensure functionality:

```bash
go test ./tests -v
```

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.