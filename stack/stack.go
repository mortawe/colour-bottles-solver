package stack

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var v T
		return v, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var v T
		return v, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element, true
}

func (s *Stack[T]) Copy() Stack[T] {
	sC := make(Stack[T], len(*s))
	copy(sC, *s)
	return sC
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Get() []T {
	return *s
}
