package utils

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) int {
	s.items = append(s.items, val)
	return len(s.items)
}

func (s *Stack[T]) PushN(vals []T) int {
	s.items = append(s.items, vals...)
	return len(s.items)
}

func (s *Stack[T]) Pop() T {
	i := len(s.items) - 1
	toRemove := s.items[i]
	s.items = s.items[:i]
	return toRemove
}

func (s *Stack[T]) PopN(n int) []T {
	toRemove := s.items[len(s.items)-n:]
	s.items = s.items[:len(s.items)-n]
	return toRemove
}

func (s *Stack[T]) Peek() T {
	i := len(s.items) - 1
	return s.items[i]
}
