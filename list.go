package socer

import (
	"fmt"
	"strings"
)

// List represents a simple generic singly linked list.
// The zero value of a List is a empty linked list ready to use.
type List[T any] struct {
	// back element of the list.
	back *ListNode[T]
	// front element of the list.
	front *ListNode[T]
	// length is the number of nodes in the list.
	length int
}

// ListNode represents a node in a linked list.
type ListNode[T any] struct {
	// value of the node.
	Value T
	// next node in the list.
	Next *ListNode[T]
}

// initListNode creates a node from value v and next node n.
func initListNode[T any](v T, n *ListNode[T]) *ListNode[T] {
	return &ListNode[T]{
		Value: v,
		Next:  n,
	}
}

// PushBack adds a element at the back of the list.
func (l *List[T]) PushBack(v T) *ListNode[T] {
	n := initListNode(v, nil)

	if l.back == nil {
		l.front = n
		l.back = n
	} else {
		l.back.Next = n
		l.back = n
	}

	l.length++

	return n
}

// PushFront adds a element at the front of the list.
func (l *List[T]) PushFront(v T) *ListNode[T] {
	n := initListNode(v, l.front)

	if l.back == nil {
		l.back = n
	}

	l.front = n
	l.length++

	return n
}

// Insert value v after n.
func (l *List[T]) InsertAfter(n *ListNode[T], v T) *ListNode[T] {
	n.Next = initListNode(v, n.Next)
	l.length++

	return n
}

// Insert value v before n. Since List is a singly linked list
// this operation will compute in O(n). This function will lead
// to undefined behavior if n isn't a node of l.
func (l *List[T]) InsertBefore(n *ListNode[T], v T) *ListNode[T] {
	for cur := l.front; cur != nil; cur = cur.Next {
		if cur.Next == n {
			cur.Next = initListNode(v, cur.Next)
			return cur.Next
		}
	}

	return nil
}

func (l *List[T]) InsertAt(i int, v T) *ListNode[T] {
	index := 0
	for n := l.front; n != nil; n = n.Next {
		if i == index {
			n.Next = initListNode(v, n.Next)
			return n.Next
		}

		index++
	}

	return nil
}

// Remove node n from the list.
func (l *List[T]) Remove(n *ListNode[T]) {
	var prev *ListNode[T]

	for cur := l.front; cur != nil; cur = cur.Next {
		if cur == n {
			prev.Next = cur.Next
		}

		prev = n
	}
}

// Back returns the back node of the list.
func (l *List[T]) Back() *ListNode[T] {
	return l.back
}

// Front returns the back node of the list.
func (l *List[T]) Front() *ListNode[T] {
	return l.front
}

// Len returns the number of nodes in the list.
func (l *List[T]) Len() int {
	return l.length
}

// IsEmpty returns true if the list is empty.
func (l *List[T]) IsEmpty() bool {
	return l.length == 0
}

// String returns a string representing the list elements.
// The string will show elements from back to front.
// Example: "[32 64 128]", where 32 is the back element.
func (l *List[T]) String() string {
	var s strings.Builder

	s.WriteRune('[')

	for n := l.front; n != nil; n = n.Next {
		fmt.Fprintf(&s, "%v", n.Value)
		if n.Next != nil {
			s.WriteRune(' ')
		}
	}

	s.WriteRune(']')

	return s.String()
}

// Iter returns a new ListIterator.
func (l *List[T]) Iter() *ListIterator[T] {
	var zero T

	return &ListIterator[T]{
		node: initListNode(zero, l.front),
	}
}

// ListIterator represents a iterator for a list.
//
// To iterate over a list (where l is a *List):
//
//	it := l.Iter()
//	for it.Next() {
//		// do something with it.Value()
//	}
type ListIterator[T any] struct {
	node *ListNode[T]
}

// Next advances the iterator to the next node.
func (it *ListIterator[T]) Next() bool {
	it.node = it.node.Next

	return it.node != nil
}

// Value returns the value of the current node.
func (it *ListIterator[T]) Value() T {
	return it.node.Value
}

// ListNode returns the current node.
func (it *ListIterator[T]) ListNode() *ListNode[T] {
	return it.node
}
