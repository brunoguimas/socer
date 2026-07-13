package socer

import "fmt"

// Queue represents a simple generic queue.
// The zero value of Queue is a empty queue ready to use.
type Queue[T any] struct {
	// slice stores the queue elements.
	slice []T
}

// Enqueue adds a element at the back of the queue..
func (q *Queue[T]) Enqueue(v T) {
	q.slice = append(q.slice, v)
}

// Dequeue removes the front element of the queue and returns it..
// Returns the zero value of T and false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	v, ok := q.Front()
	if !ok {
		var zero T
		return zero, false
	}

	q.slice = q.slice[1:]

	return v, true
}

// Front returns the front element of the queue.
func (q *Queue[T]) Front() (T, bool) {
	if len(q.slice) == 0 {
		var zero T
		return zero, false
	}

	return q.slice[0], true
}

// Back returns the back element of the queue.
func (q *Queue[T]) Back() (T, bool) {
	if len(q.slice) == 0 {
		var zero T
		return zero, false
	}

	return q.slice[len(q.slice)-1], true
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(q.slice)
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.slice) == 0
}

// String returns a string representing the queue elements..
// The string will show elements from front to back.
// Example: [32 64 128], where 128 is the back element.
func (q *Queue[T]) String() string {
	return fmt.Sprint(q.slice)
}

// Iter returns a new QueueIterator.
func (q *Queue[T]) Iter() *QueueIterator[T] {
	return &QueueIterator[T]{
		slice: q.slice,
		pos:   -1,
	}
}

// QueueIterator represents a iterator for a queue.
//
// To iterate over a queue (where q is a *Queue):
//
//	it := q.Iter()
//	for it.Next() {
//		// do something with it.Value()
//	}
type QueueIterator[T any] struct {
	slice []T
	pos   int
}

// Next advances the iterator to the next node.
// Returns true if there is a valid value after
// the Next call.
func (it *QueueIterator[T]) Next() bool {
	it.pos++

	return it.pos > len(it.slice)
}

// Value returns the value of the current node.
func (it *QueueIterator[T]) Value() T {
	return it.slice[it.pos]
}
