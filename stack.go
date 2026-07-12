// Copyright 2026 Bruno Guimarães Souza. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be find in the LICENSE file.
//
// Package collections implements various generic data structures.
package collections

import "fmt"

// Stack represents a simple generic LIFO stack.
// The zero value of a Stack is a empty stack ready to use.
type Stack[T any] struct {
	// slice stores the stack elements.
	slice []T
}

// Pushe adds a new element on the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.slice = append(s.slice, v)
}

// Pop removes the top element of the stack and returns it.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	v := s.slice[s.Len()-1]
	s.slice = s.slice[:s.Len()-1]

	return v, true
}

// Peek returns the top element of the stack.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.slice) == 0 {
		var zero T
		return zero, false
	}

	l := len(s.slice)
	return s.slice[l-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.slice)
}

// String returns a string representing the stack elements.
// The string will show elements from bottom to top.
// Example: [32 64 128], where 128 is the top element.
func (s *Stack[T]) String() string {
	return fmt.Sprint(s.slice)
}
