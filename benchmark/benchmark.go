package benchmark

import "sync"

type Benchmark struct {
	mu       sync.Mutex
	hits     int
	misses   int
	expired  int
}

func NewBenchmark() *Benchmark {
	return &Benchmark{}
}

func (b *Benchmark) RecordHit() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.hits++
}

func (b *Benchmark) RecordMiss() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.misses++
}

func (b *Benchmark) RecordExpiration() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.expired++
}

func (b *Benchmark) Hits() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.hits
}

func (b *Benchmark) Misses() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.misses
}

func (b *Benchmark) Expired() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.expired
}
