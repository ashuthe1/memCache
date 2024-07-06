package cache

import (
	"testing"
	"time"

	"github.com/ashuthe1/kuki-memcache/eviction"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := NewCache(2*time.Second, 2, eviction.NewLRU(), nil)

	c.Set("key1", "value1")
	c.Set("key2", "value2")

	val, found := c.Get("key1")
	assert.True(t, found)
	assert.Equal(t, "value1", val)

	time.Sleep(3 * time.Second)

	_, found = c.Get("key1")
	assert.False(t, found)

	c.Set("key3", "value3")
	_, found = c.Get("key2")
	assert.False(t, found)
}
