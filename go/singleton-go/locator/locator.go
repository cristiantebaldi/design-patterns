package locator

import (
	"fmt"
	"sync"
)

type ServiceLocator struct {
	mu       sync.RWMutex
	services map[string]any
}

var (
	instance *ServiceLocator
	once     sync.Once
)

func GetInstance() *ServiceLocator {
	once.Do(func() {
		instance = &ServiceLocator{services: make(map[string]any)}
	})
	return instance
}

func Register[T any](l *ServiceLocator, name string, service T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.services[name] = service
}

func Get[T any](l *ServiceLocator, name string) (T, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var zero T
	svc, ok := l.services[name]
	if !ok {
		return zero, fmt.Errorf("service %s not found", name)
	}
	typed, ok := svc.(T)
	if !ok {
		return zero, fmt.Errorf("service %s has wrong type", name)
	}
	return typed, nil
}