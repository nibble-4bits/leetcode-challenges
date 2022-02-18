package main

import "errors"

// runeStack represents a stack data structure that can hold runes.
type runeStack struct {
	items []rune
	count int
}

// push places a new rune at the top of the stack.
func (s *runeStack) push(r rune) {
	s.items = append(s.items, r)
	s.count++
}

// pop removes the rune at the top of the stack and returns it.
func (s *runeStack) pop() (rune, error) {
	if s.count == 0 {
		return 0, errors.New("cannot pop item of empty stack")
	}

	val := s.items[s.count-1]
	s.items = s.items[0 : s.count-1]
	s.count--

	return val, nil
}
