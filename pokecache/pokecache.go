package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries      map[string]cacheEntry
	mu           sync.RWMutex
	reapTicker   *time.Ticker
	reapInterval time.Duration
}

func NewCache(reapInterval time.Duration) *Cache {
	c := Cache{
		entries:      make(map[string]cacheEntry),
		mu:           sync.RWMutex{},
		reapTicker:   time.NewTicker(reapInterval),
		reapInterval: reapInterval,
	}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	for range c.reapTicker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > c.reapInterval {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) Close() {
	c.reapTicker.Stop()
}
