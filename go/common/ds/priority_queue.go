package ds

type MinPriorityQueue[T any] struct {
	heap *MinHeap[T]
}

func NewMinPriorityQueue[T any](getPriorityFn func(item T) int) *MinPriorityQueue[T] {
	return &MinPriorityQueue[T]{
		heap: NewMinHeap[T](getPriorityFn),
	}
}

func (pq *MinPriorityQueue[T]) Enqueue(value T) {
	pq.heap.Insert(value)
}

func (pq *MinPriorityQueue[T]) Dequeue() (T, error) {
	if pq.heap.IsEmpty() {
		return pq.getNull(), NewQueueEmptyException()
	}

	return pq.heap.ExtractMin(), nil
}

func (pq *MinPriorityQueue[T]) IsEmpty() bool {
	return pq.heap.IsEmpty()
}

func (pq *MinPriorityQueue[T]) Peek() (T, error) {
	if pq.heap.IsEmpty() {
		return pq.getNull(), NewQueueEmptyException()
	}

	return pq.heap.GetMin(), nil
}

func (pq *MinPriorityQueue[T]) getNull() T {
	var null T

	return null
}
