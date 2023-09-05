package store

import (
	"sync"
)

type ConcurrentQueue[T any] struct {
	mu   sync.Mutex
	data []T
}

func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{
		data: make([]T, 0),
	}
}

// Enqueue adds an element to the end of the queue
func (cq *ConcurrentQueue[T]) Enqueue(value T) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.data = append(cq.data, value)
}

// Dequeue removes an element from the start of the queue
func (cq *ConcurrentQueue[T]) Dequeue() (T, bool) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.data) == 0 {
		var zero T
		return zero, false
	}

	val := cq.data[0]
	cq.data = cq.data[1:]

	return val, true
}

// Peek returns the first element in the queue without removing it
func (cq *ConcurrentQueue[T]) Peek() (T, bool) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.data) == 0 {
		var zero T
		return zero, false
	}

	return cq.data[0], true
}

func (cq *ConcurrentQueue[T]) Len() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.data)
}
