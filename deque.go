package socer

import "fmt"

// Deque represents a simple generic deque.
// The zero value of Deque is a empty deque ready to use.
type Deque[T any] struct {
	// slice stores the deque elements.
	slice []T
}

// PushFront adds a element at the front of the deque.
func (d *Deque[T]) PushFront(v T) {
	d.slice = append(d.slice, v)
}

// PushBack adds a element at the back of the deque.
func (d *Deque[T]) PushBack(v T) {
	d.slice = append(d.slice[:0], append([]T{v}, d.slice[0:]...)...)
}

// PopFront removes the front element of the deque and returns it..
// Returns the zero value of T and false if the deque is empty.
func (d *Deque[T]) PopFront() (T, bool) {
	v, ok := d.Front()
	if !ok {
		var zero T
		return zero, false
	}

	d.slice = d.slice[1:]

	return v, true
}

// PopBack removes the back element of the deque and returns it..
// Returns the zero value of T and false if the deque is empty.
func (d *Deque[T]) PopBack() (T, bool) {
	v, ok := d.Back()
	if !ok {
		var zero T
		return zero, false
	}

	d.slice = d.slice[:d.Len()-1]
	return v, true
}

// Front returns the front element of the deque.
func (d *Deque[T]) Front() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	return d.slice[0], true
}

// Back returns the back element of the deque.
func (d *Deque[T]) Back() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	return d.slice[d.Len()-1], true
}

// Len returns the number of elements in the deque.
func (d *Deque[T]) Len() int {
	return len(d.slice)
}

// IsEmpty returns true if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.Len() == 0
}

// String returns a string representing the deque elements.
// The string will show elements from front to back.
// Example: "[128 64 32]", where 128 is the front element.
func (d *Deque[T]) String() string {
	return fmt.Sprint(d.slice)
}

// Iter returns a new DequeIterator.
func (q *Deque[T]) Iter() *DequeIterator[T] {
	return &DequeIterator[T]{
		slice: q.slice,
		pos:   -1,
	}
}

// DequeIterator represents a iterator for a deque.
//
// To iterate over a deque (where d is a *Deque):
//
//	it := d.Iter()
//	for it.Next() {
//		// do something with it.Value()
//	}
type DequeIterator[T any] struct {
	slice []T
	pos   int
}

// Next advances the iterator to the next node.
// Returns true if there is a valid value after
// the Next call.
func (it *DequeIterator[T]) Next() bool {
	it.pos++

	return it.pos > len(it.slice)
}

// Value returns the value of the current node.
func (it *DequeIterator[T]) Value() T {
	return it.slice[it.pos]
}
