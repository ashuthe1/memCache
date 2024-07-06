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
	mu        sync.Mutex
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
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		c.benchmark.RecordMiss()
		return nil, false
	}

	if time.Now().After(item.ExpiresAt) {
		delete(c.items, key)
		c.policy.Remove(key)
		c.benchmark.RecordExpiration()
		if c.onEvicted != nil {
			c.onEvicted(key, item.Value)
		}
		return nil, false
	}

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
	return c.benchmark.Hits()
}

func (c *Cache) Misses() int {
	return c.benchmark.Misses()
}

func (c *Cache) Expired() int {
	return c.benchmark.Expired()
}
