package eviction

import "container/list"

type FIFO struct {
	queue *list.List
}

func NewFIFO() *FIFO {
	return &FIFO{
		queue: list.New(),
	}
}

func (f *FIFO) Add(key string) {
	f.queue.PushBack(key)
}

func (f *FIFO) Remove(key string) {
	for e := f.queue.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			f.queue.Remove(e)
			break
		}
	}
}

func (f *FIFO) Evict() string {
	front := f.queue.Front()
	if front != nil {
		key := front.Value.(string)
		f.queue.Remove(front)
		return key
	}
	return ""
}
