package cache

import "sync"

type Cache interface {
	Set(key string, value any)
	Get(key string) (any, bool)
	Has(key string) bool
	Clear()
}

type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]any
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string]any)}
}

func (c *MemoryCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *MemoryCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *MemoryCache) Has(key string) bool {
	_, ok := c.Get(key)
	return ok
}

func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]any)
}