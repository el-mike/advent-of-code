package ds

type MaxHeap[T any] struct {
	BaseHeap[T]

	moreFn func(data []T, i, j int) bool
}

func NewMaxHeap[T any](moreFn func(data []T, i, j int) bool) *MaxHeap[T] {
	return &MaxHeap[T]{
		BaseHeap: BaseHeap[T]{
			Data: []T{},
		},
		moreFn: moreFn,
	}
}

func (mh *MaxHeap[T]) Build(data []T) {
	tmp := make([]T, len(data))
	copy(tmp, data)

	mh.Data = tmp

	mh.heapify()
}

func (mh *MaxHeap[T]) GetMax() T {
	if len(mh.Data) == 0 {
		return mh.getNull()
	}

	return mh.Data[0]
}

func (mh *MaxHeap[T]) ExtractMax() T {
	l := len(mh.Data)

	if l == 0 {
		return mh.getNull()
	}

	min := mh.GetMax()

	mh.Data[0], mh.Data[l-1] = mh.Data[l-1], mh.Data[0]
	mh.Data = mh.Data[:(l - 1)]

	mh.siftDown(0)

	return min
}

func (mh *MaxHeap[T]) Insert(value T) {
	mh.Data = append(mh.Data, value)
	l := len(mh.Data)

	if l == 1 {
		return
	}

	// Since new element is always added to the end of the underlying slice, we need to sift it
	// up to find the correct place.
	mh.siftUp(l - 1)
}

func (mh *MaxHeap[T]) siftUp(currentIndex int) {
	parentIndex := mh.getParent(currentIndex)

	for currentIndex != 0 && mh.moreFn(mh.Data, currentIndex, parentIndex) {
		mh.Data[currentIndex], mh.Data[parentIndex] = mh.Data[parentIndex], mh.Data[currentIndex]

		currentIndex = parentIndex
		parentIndex = mh.getParent(currentIndex)
	}
}

func (mh *MaxHeap[T]) siftDown(currentIndex int) {
	leftIndex := mh.getLeftChildIndex(currentIndex)
	// Right child index is always right next to the left child.
	rightIndex := leftIndex + 1

	l := len(mh.Data)

	// First condition stops sifting down for the last level - when currentIndex is in the last level,
	// left and right indices will always be out of bounds.
	for (leftIndex < l && !mh.moreFn(mh.Data, currentIndex, leftIndex)) ||
		(rightIndex < l && !mh.moreFn(mh.Data, currentIndex, rightIndex)) {
		maxIndex := 0

		if rightIndex >= l || mh.moreFn(mh.Data, leftIndex, rightIndex) {
			maxIndex = leftIndex
		} else {
			maxIndex = rightIndex
		}

		mh.Data[currentIndex], mh.Data[maxIndex] = mh.Data[maxIndex], mh.Data[currentIndex]

		currentIndex = maxIndex
		leftIndex = mh.getLeftChildIndex(currentIndex)
		rightIndex = leftIndex + 1
	}
}

// heapify - ensures that heap property is preserved. It uses sifting down, as it's
// faster than sifting up (the lower levels have more elements)
func (mh *MaxHeap[T]) heapify() {
	// We could start i at (size / 2) - 1, as the last level cannot be sifted down,
	// but it's taken care of automatically in siftDown method (loop's condition).
	for i := len(mh.Data) - 1; i >= 0; i -= 1 {
		mh.siftDown(i)
	}
}
