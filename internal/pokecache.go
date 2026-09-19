package internal

import (
	"sync"
	"time"
)

type pokeCash struct {
	Entries map[string]cacheEntry
	mu      sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *pokeCash {
	c := &pokeCash{
		Entries: make(map[string]cacheEntry),
	}
	go c.readLoop(interval)
	return c
}

func (c *pokeCash) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{val: val, createdAt: time.Now()}
}

func (c *pokeCash) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cac *pokeCash) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(-interval)

		cac.mu.Lock()
		for key, entry := range cac.Entries {
			if entry.createdAt.Before(cutoff) {
				delete(cac.Entries, key)
			}
		}
		cac.mu.Unlock()
	}
}
