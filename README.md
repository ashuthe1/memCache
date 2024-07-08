# In-Memory Caching Library

This project implements an in-memory caching library in Golang with support for multiple eviction policies (FIFO, LRU, LIFO) and custom eviction policies.

## Features

- **Eviction Policies**: Supports FIFO, LRU, and LIFO eviction policies.
- **Custom Eviction Policy**: Allows users to define and integrate custom eviction policies.
- **Thread Safety**: Utilizes `sync.RWMutex` for concurrent read/write access safety.
- **Statistics**: Tracks cache hits, misses, and expired items.
- **Benchmarking**: Basic benchmarking capabilities for cache operations.
- **Versatile Value Storage**: Can store any type of value, including integers, strings, arrays, maps, and structs.
- **Batch Operations**: Supports setting and getting multiple key-value pairs at once.
- **Persistence**: Save cache contents to a file and load from a file.

### Folder Structure
```
kuki-memcache/
├── cache/
│   └── cache.go
├── benchmark/
│   └── benchmark.go
├── eviction/
│   ├── policy.go
│   ├── fifo.go
│   ├── lifo.go
│   └── lru.go
├── persistence/
│   ├── file_persistence.go
│   └── persistable.go
├── test/
│   └── cache_test.go
├── sample/
│   └── main.go
└── README.md
```

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
	// Example: Create a cache with 20 minutes TTL, size of 10 items, and LRU eviction policy
	c := cache.NewCache(20*time.Minute, 10, eviction.NewLRU(), nil)

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
```

### Batch Operations

You can add and retrieve multiple key-value pairs at once using the `BatchSet` and `BatchGet` methods.

```go
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
```

### Supported Operations

- **NewCache()**: Create a new cache instance with TTL, maximum size, and an eviction policy.
  - **Parameters**:
    - `ttl` (time.Duration): The time-to-live duration for each cache entry.
    - `maxSize` (int): The maximum number of items the cache can hold.
    - `evictionPolicy` (EvictionPolicy): The eviction policy to use (e.g., FIFO, LRU, LIFO, or a custom policy).
    - `logger` (optional): A logger instance for logging cache operations (can be nil).
  - **Return Value**: A new cache instance.

- **Set()**: Add a key-value pair to the cache.
  - **Parameters**:
    - `key` (string): The key for the cache entry.
    - `value` (interface{}): The value to be stored in the cache.
  - **Return Value**: None.

- **Get()**: Retrieve a value associated with a key from the cache.
  - **Parameters**:
    - `key` (string): The key for the cache entry.
  - **Return Value**: A tuple containing:
    - `value` (interface{}): The value associated with the key, or `nil` if not found.
    - `isExists` (bool): A boolean indicating whether the key was found in the cache.

- **BatchSet()**: Add multiple key-value pairs to the cache.
  - **Parameters**:
    - `items` (map[string]interface{}): The key-value pairs to be stored in the cache.
  - **Return Value**: None.

- **BatchGet()**: Retrieve multiple key-value pairs from the cache.
  - **Parameters**:
    - `keys` ([]string): The keys for the cache entries to be retrieved.
  - **Return Value**: A map containing the key-value pairs that were found in the cache.

- **Delete()**: Remove a key-value pair from the cache.
  - **Parameters**:
    - `key` (string): The key for the cache entry.
  - **Return Value**: None.

- **Hits()**: Track cache hits.
  - **Parameters**: None.
  - **Return Value**: An integer representing the number of cache hits.

- **Misses()**: Track cache misses.
  - **Parameters**: None.
  - **Return Value**: An integer representing the number of cache misses.

- **Expired()**: Track expired items.
  - **Parameters**: None.
  - **Return Value**: An integer representing the number of expired cache items.

### Running Tests

Run tests to ensure functionality:

```bash
go test ./tests -v
```

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.