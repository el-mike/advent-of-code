package common

import "errors"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.data) == 0 {
		return q.getNull(), errors.New("Queue is empty")
	}

	value := q.data[0]
	q.data = q.data[1:]

	return value, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.data) == 0 {
		return q.getNull(), errors.New("Queue is empty")
	}

	return q.data[0], nil
}

func (q *Queue[T]) getNull() T {
	var null T

	return null
}
