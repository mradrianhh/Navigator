package models

import (
	"errors"
)

// NavigationStack ..
type NavigationStack []Screen

// IsEmpty checks if the stack is empty.
func (s *NavigationStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack.
func (s *NavigationStack) Push(screen Screen) {
	*s = append(*s, screen)
}

// Pop removes and returns the top element of the stack.
func (s *NavigationStack) Pop() (Screen, error) {
	if s.IsEmpty() {
		return Empty{}, errors.New("stack is empty")
	}

	index := len(*s) - 1
	screen := (*s)[index]
	*s = (*s)[:index]
	return screen, nil
}

// Peek returns the top element of the stack without removing it.
func (s *NavigationStack) Peek() (Screen, error) {
	if s.IsEmpty() {
		return Empty{}, errors.New("stack is empty")
	}

	index := len(*s) - 1
	screen := (*s)[index]
	return screen, nil
}

// Len returns the length of the stack.
func (s *NavigationStack) Len() int {
	return len(*s)
}
