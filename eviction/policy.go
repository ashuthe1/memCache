package eviction

import (
	"fmt"
)

type EvictionPolicy interface {
	Add(key string)
	Remove(key string)
	Evict() string
}

func NewEvictionPolicy(policyType interface{}) EvictionPolicy {
	switch policyType.(type) {
	case FIFO:
		return NewFIFO()
	case LRU:
		return NewLRU()
	case LIFO:
		return NewLIFO()
	default:
		fmt.Println("Invalid eviction policy type, using FIFO as default")
		return NewFIFO()
	}
}
