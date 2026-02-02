package queue

// Queue represents a generic FIFO (First In First Out) queue.
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an element to the rear of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

// Dequeue removes and returns the front element.
// ok is false if the queue is empty.
func (q *Queue[T]) Dequeue() (v T, ok bool) {
	if len(q.items) == 0 {
		return v, false
	}
	v = q.items[0]
	q.items = q.items[1:]
	return v, true
}

// Front returns the front element without removing it.
// ok is false if the queue is empty.
func (q *Queue[T]) Front() (v T, ok bool) {
	if len(q.items) == 0 {
		return v, false
	}
	return q.items[0], true
}

// Rear returns the rear element without removing it.
// ok is false if the queue is empty.
func (q *Queue[T]) Rear() (v T, ok bool) {
	if len(q.items) == 0 {
		return v, false
	}
	return q.items[len(q.items)-1], true
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.items = make([]T, 0)
}

// ToSlice returns a slice containing all elements in the queue (from front to rear).
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, len(q.items))
	copy(result, q.items)
	return result
}
