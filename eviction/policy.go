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
	case CustomEviction:
		return NewCustom()
	default:
		panic(fmt.Sprintf("unsupported eviction policy type: %v", policyType))
	}
}
