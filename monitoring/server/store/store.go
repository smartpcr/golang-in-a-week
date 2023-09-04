package store

import "sync"

// ConcurrentMap is a simple generic concurrent map using a regular map and a mutex.
type ConcurrentMap[K comparable, V any] struct {
	m    map[K]V
	lock sync.Mutex
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		m: make(map[K]V),
	}
}

func (cm *ConcurrentMap[K, V]) Store(key K, value V) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.m[key] = value
}

func (cm *ConcurrentMap[K, V]) Find(key K) (V, bool) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	v, ok := cm.m[key]
	return v, ok
}

func (cm *ConcurrentMap[K, V]) Delete(key K) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	delete(cm.m, key)
}
