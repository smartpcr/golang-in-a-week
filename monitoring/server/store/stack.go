package store

import "sync"

type ConcurrentStack[T any] struct {
	mu   sync.Mutex
	data []T
}

func NewConcurrentStack[T any]() *ConcurrentStack[T] {
	return &ConcurrentStack[T]{
		data: make([]T, 0),
	}
}

func (cs *ConcurrentStack[T]) Push(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.data = append(cs.data, value)
}

func (cs *ConcurrentStack[T]) Pop() (T, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if len(cs.data) == 0 {
		var zero T
		return zero, false
	}

	val := cs.data[len(cs.data)-1]
	cs.data = cs.data[:len(cs.data)-1]

	return val, true
}

func (cs *ConcurrentStack[T]) Peek() (T, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if len(cs.data) == 0 {
		var zero T
		return zero, false
	}

	return cs.data[len(cs.data)-1], true
}

func (cs *ConcurrentStack[T]) Len() int {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return len(cs.data)
}
