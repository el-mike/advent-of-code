package ds

type BaseHeap[T any] struct {
	Data []T
}

func (bh *BaseHeap[T]) IsEmpty() bool {
	return len(bh.Data) == 0
}

func (bh *BaseHeap[T]) getParent(i int) int {
	// We use zero-based array for the heap, therefore we need to subtract one
	// to get the correct index.
	return (i - 1) / 2
}

func (bh *BaseHeap[T]) getLeftChildIndex(i int) int {
	// We use zero-based array for the heap, therefore we need to add one
	// to get the correct index.
	return (2 * i) + 1
}

func (bh *BaseHeap[T]) getNull() T {
	var null T

	return null
}
