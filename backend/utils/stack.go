package utils

type Stack[T any] struct {
	elements []T
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() *T {
	if !s.IsEmpty() {
		element := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return &element
	}
	return nil
}

func (s *Stack[T]) Top() *T {
	return &s.elements[len(s.elements)-1]
}