package ds

type HeapType uint8

const (
	MinHeapType HeapType = iota
	MaxHeapType
)

type BaseHeap[T any] struct {
	Data []T
	Type HeapType

	getScoreFn func(item T) int
}

func (bh *BaseHeap[T]) Insert(value T) {
	bh.Data = append(bh.Data, value)
	l := len(bh.Data)

	if l == 1 {
		return
	}

	// Since new element is always added to the end of the underlying slice, we need to sift it
	// up to find the correct place.
	bh.siftUp(l - 1)
}

func (bh *BaseHeap[T]) IsEmpty() bool {
	return len(bh.Data) == 0
}

func (mh *MinHeap[T]) Build(data []T) {
	tmp := make([]T, len(data))
	copy(tmp, data)

	mh.Data = tmp

	mh.heapify()
}

func (bh *BaseHeap[T]) siftUp(currentIndex int) {
	parentIndex := bh.getParent(currentIndex)

	for currentIndex != 0 && bh.compare(currentIndex, parentIndex) {
		bh.Data[currentIndex], bh.Data[parentIndex] = bh.Data[parentIndex], bh.Data[currentIndex]

		currentIndex = parentIndex
		parentIndex = bh.getParent(currentIndex)
	}
}

func (bh *BaseHeap[T]) siftDown(currentIndex int) {
	leftIndex := bh.getLeftChildIndex(currentIndex)
	// Right child index is always right next to the left child.
	rightIndex := leftIndex + 1

	l := len(bh.Data)

	// First condition stops sifting down for the last level - when currentIndex is in the last level,
	// left and right indices will always be out of bounds.
	for (leftIndex < l && !bh.compare(currentIndex, leftIndex)) ||
		(rightIndex < l && !bh.compare(currentIndex, rightIndex)) {
		maxIndex := 0

		if rightIndex >= l || bh.compare(leftIndex, rightIndex) {
			maxIndex = leftIndex
		} else {
			maxIndex = rightIndex
		}

		bh.Data[currentIndex], bh.Data[maxIndex] = bh.Data[maxIndex], bh.Data[currentIndex]

		currentIndex = maxIndex
		leftIndex = bh.getLeftChildIndex(currentIndex)
		rightIndex = leftIndex + 1
	}
}

// heapify - ensures that heap property is preserved. It uses sifting down, as it's
// faster than sifting up (the lower levels have more elements)
func (mh *MinHeap[T]) heapify() {
	// We could start i at (size / 2) - 1, as the last level cannot be sifted down,
	// but it's taken care of automatically in siftDown method (loop's condition).
	for i := len(mh.Data) - 1; i >= 0; i -= 1 {
		mh.siftDown(i)
	}
}

// compare - compares two elements and decides whether they should be swapped or not.
func (bh *BaseHeap[T]) compare(i, j int) bool {
	iScore, jScore := bh.getScoreFn(bh.Data[i]), bh.getScoreFn(bh.Data[j])

	if bh.Type == MinHeapType {
		return iScore < jScore
	} else {
		return iScore > jScore
	}
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
