package cache

import "sync"

type SingletonCache struct {
	mu    sync.RWMutex
	items map[string]any
}

var (
	instance *SingletonCache
	once     sync.Once
)

func GetInstance() *SingletonCache {
	once.Do(func() {
		instance = &SingletonCache{
			items: make(map[string]any),
		}
	})
	return instance
}

func (c *SingletonCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *SingletonCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[key]
	return v, ok
}

func (c *SingletonCache) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.items[key]
	return ok
}

func (c *SingletonCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]any)
}