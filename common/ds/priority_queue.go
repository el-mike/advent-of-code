package ds

type PriorityQueue[T any] struct {
	heap *MaxHeap[T]
}

func NewPriorityQueue[T any](moreFn func(data []T, i, j int) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: NewMaxHeap[T](moreFn),
	}
}

func (pq *PriorityQueue[T]) Enqueue(value T) {
	pq.heap.Insert(value)
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	if pq.heap.IsEmpty() {
		return pq.getNull(), NewQueueEmptyException()
	}

	return pq.heap.ExtractMax(), nil
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.heap.IsEmpty()
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.heap.IsEmpty() {
		return pq.getNull(), NewQueueEmptyException()
	}

	return pq.heap.GetMax(), nil
}

func (pq *PriorityQueue[T]) getNull() T {
	var null T

	return null
}
