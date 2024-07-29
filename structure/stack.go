package structure

import (
	"sync"
)

type Stack[T any] struct {
	lock   sync.Mutex
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0)}
}
func (s *Stack[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.values = append(s.values, v)
}
func (s *Stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	l := len(s.values)
	if l == 0 {
		var empty T
		return empty, ErrStackEmpty
	}
	res := s.values[l-1]
	s.values = s.values[:l-1]
	return res, nil
}
func (stack *Stack[T]) Peek() (T, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	l := len(stack.values)
	if l == 0 {
		var empty T
		return empty, ErrStackEmpty
	}
	return stack.values[l-1], nil
}
func (stack *Stack[T]) Clear() {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.values = nil
}
