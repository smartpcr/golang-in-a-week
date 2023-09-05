package store

import "sync"

type Node[T comparable] struct {
	mu       sync.RWMutex
	Value    T
	children []*Node[T]
}

type ConcurrentTree[T comparable] struct {
	Root *Node[T]
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value:    value,
		children: make([]*Node[T], 0),
	}
}

func (n *Node[T]) AddChild(child *Node[T]) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.children = append(n.children, child)
}

func (n *Node[T]) GetChildren() []*Node[T] {
	n.mu.RLock()
	defer n.mu.RUnlock()

	// Create a shallow copy of children slice for safety
	childrenCopy := make([]*Node[T], len(n.children))
	copy(childrenCopy, n.children)
	return childrenCopy
}

func NewConcurrentTree[T comparable](rootValue T) *ConcurrentTree[T] {
	return &ConcurrentTree[T]{
		Root: NewNode(rootValue),
	}
}

func (n *Node[T]) Count() int {
	n.mu.RLock()
	defer n.mu.RUnlock()
	count := 1
	for _, child := range n.children {
		count += child.Count()
	}
	return count
}

func (n *Node[T]) Depth() int {
	n.mu.RLock()
	defer n.mu.RUnlock()
	maxDepth := 1
	for _, child := range n.children {
		childDepth := child.Depth()
		if childDepth+1 > maxDepth {
			maxDepth = childDepth + 1
		}
	}
	return maxDepth
}

func (n *Node[T]) Find(predicate func(T) bool) *Node[T] {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if predicate(n.Value) {
		return n
	}
	for _, child := range n.children {
		found := child.Find(predicate)
		if found != nil {
			return found
		}
	}
	return nil
}

func (n *Node[T]) FindAll(predicate func(T) bool) []*Node[T] {
	n.mu.RLock()
	defer n.mu.RUnlock()

	var results []*Node[T]

	if predicate(n.Value) {
		results = append(results, n)
	}

	for _, child := range n.children {
		results = append(results, child.FindAll(predicate)...)
	}

	return results
}

func (n *Node[T]) AllUniqueValue(getValue func(T) string) []string {
	n.mu.RLock()
	defer n.mu.RUnlock()

	uniqueValues := make(map[string]struct{}) // This will act as our HashSet

	var collect func(node *Node[T])
	collect = func(node *Node[T]) {
		if val := getValue(node.Value); val != "" {
			uniqueValues[val] = struct{}{}
		}

		for _, child := range node.children {
			collect(child)
		}
	}

	collect(n)

	values := make([]string, 0, len(uniqueValues))
	for key := range uniqueValues {
		values = append(values, key)
	}

	return values
}
