package ds

type StackEmptyException struct{}

func NewStackEmptyException() *StackEmptyException {
	return &StackEmptyException{}
}

func (e *StackEmptyException) Error() string {
	return "Stack is empty"
}

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() (T, error) {
	l := len(s.data)

	if len(s.data) == 0 {
		return s.getNull(), NewStackEmptyException()
	}

	value := s.data[l-1]
	s.data = s.data[:(l - 1)]

	return value, nil
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Peek() (T, error) {
	l := len(s.data)

	if l == 0 {
		return s.getNull(), NewStackEmptyException()
	}

	return s.data[l-1], nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) getNull() T {
	var null T

	return null
}
