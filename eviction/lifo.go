package eviction

import "container/list"

type LIFO struct {
	stack *list.List
}

func NewLIFO() *LIFO {
	return &LIFO{
		stack: list.New(),
	}
}

func (l *LIFO) Add(key string) {
	l.stack.PushBack(key)
}

func (l *LIFO) Remove(key string) {
	for e := l.stack.Back(); e != nil; e = e.Prev() {
		if e.Value == key {
			l.stack.Remove(e)
			break
		}
	}
}

func (l *LIFO) Evict() string {
	back := l.stack.Back()
	if back != nil {
		key := back.Value.(string)
		l.stack.Remove(back)
		return key
	}
	return ""
}
