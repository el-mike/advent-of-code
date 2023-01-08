package ds

type MinHeap[T any] struct {
	BaseHeap[T]
}

func NewMinHeap[T any](getScoreFn func(item T) int) *MinHeap[T] {
	return &MinHeap[T]{
		BaseHeap: BaseHeap[T]{
			Data:       []T{},
			Type:       MinHeapType,
			getScoreFn: getScoreFn,
		},
	}
}

func (mh *MinHeap[T]) GetMin() T {
	if len(mh.Data) == 0 {
		return mh.getNull()
	}

	return mh.Data[0]
}

func (mh *MinHeap[T]) ExtractMin() T {
	l := len(mh.Data)

	if l == 0 {
		return mh.getNull()
	}

	min := mh.GetMin()

	mh.Data[0], mh.Data[l-1] = mh.Data[l-1], mh.Data[0]
	mh.Data = mh.Data[:(l - 1)]

	mh.siftDown(0)

	return min
}
