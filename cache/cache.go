package cache

import (
	"sync"
	"time"

	"github.com/ashuthe1/kuki-memcache/benchmark"
	"github.com/ashuthe1/kuki-memcache/eviction"
)

type CacheItem struct {
	Value     interface{}
	ExpiresAt time.Time
}

type Cache struct {
	mu        sync.RWMutex
	items     map[string]CacheItem
	ttl       time.Duration
	maxSize   int
	policy    eviction.EvictionPolicy
	onEvicted func(string, interface{})
	benchmark *benchmark.Benchmark
}

func NewCache(ttl time.Duration, maxSize int, policy eviction.EvictionPolicy, onEvicted func(string, interface{})) *Cache {
	return &Cache{
		items:     make(map[string]CacheItem),
		ttl:       ttl,
		maxSize:   maxSize,
		policy:    policy,
		onEvicted: onEvicted,
		benchmark: benchmark.NewBenchmark(),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.items) >= c.maxSize {
		evictedKey := c.policy.Evict()
		if item, found := c.items[evictedKey]; found {
			delete(c.items, evictedKey)
			c.benchmark.RecordExpiration()
			if c.onEvicted != nil {
				c.onEvicted(evictedKey, item.Value)
			}
		}
	}

	c.items[key] = CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
	}
	c.policy.Add(key)
}
func (c *Cache) Get(key string) (interface{}, bool) {
    // Initial read lock to safely access the cache
    c.mu.RLock()
    // Ensure the read lock is released when the function exits
    defer c.mu.RUnlock()

    // Check if the item exists in the cache
    item, found := c.items[key]
    if !found {
        // Record a cache miss if the item is not found
        c.benchmark.RecordMiss()
        return nil, false
    }

    // Check if the item has expired
    if time.Now().After(item.ExpiresAt) {
        // Upgrade to a write lock to modify the cache
        c.mu.RUnlock()
        c.mu.Lock()
        // Double-check if the item is still expired after acquiring the write lock
        item, found = c.items[key]
        if found && time.Now().After(item.ExpiresAt) {
            // Remove the expired item from the cache
            delete(c.items, key)
            c.policy.Remove(key)
            c.benchmark.RecordExpiration()
            if c.onEvicted != nil {
                c.onEvicted(key, item.Value)
            }
        }
        // Release the write lock and downgrade to a read lock
        c.mu.Unlock()
        c.mu.RLock()

        return nil, false
    }

    // Record a cache hit if the item is found and not expired
    c.benchmark.RecordHit()
    return item.Value, true
}


func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.items[key]; found {
		delete(c.items, key)
		c.policy.Remove(key)
		if c.onEvicted != nil {
			c.onEvicted(key, item.Value)
		}
	}
}

func (c *Cache) Hits() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.benchmark.Hits()
}

func (c *Cache) Misses() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.benchmark.Misses()
}

func (c *Cache) Expired() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.benchmark.Expired()
}
