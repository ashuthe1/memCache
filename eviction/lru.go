package eviction

import "container/list"

type LRU struct {
	queue     *list.List
	cacheKeys map[string]*list.Element
}

func NewLRU() *LRU {
	return &LRU{
		queue:     list.New(),
		cacheKeys: make(map[string]*list.Element),
	}
}

func (l *LRU) Add(key string) {
	if el, found := l.cacheKeys[key]; found {
		l.queue.MoveToFront(el)
		return
	}
	el := l.queue.PushFront(key)
	l.cacheKeys[key] = el
}

func (l *LRU) Remove(key string) {
	if el, found := l.cacheKeys[key]; found {
		l.queue.Remove(el)
		delete(l.cacheKeys, key)
	}
}

func (l *LRU) Evict() string {
	back := l.queue.Back()
	if back != nil {
		key := back.Value.(string)
		l.queue.Remove(back)
		delete(l.cacheKeys, key)
		return key
	}
	return ""
}
