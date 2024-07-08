package persistence

import (
	"github.com/ashuthe1/kuki-memcache/cache"
)

// CachePersistable defines methods that must be implemented by a cache for persistence.
type CachePersistable interface {
	Lock()
	Unlock()
	Items() map[string]cache.CacheItem
}
