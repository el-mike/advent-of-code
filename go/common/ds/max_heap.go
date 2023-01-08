package ds

type MaxHeap[T any] struct {
	BaseHeap[T]
}

func NewMaxHeap[T any](getScoreFn func(item T) int) *MaxHeap[T] {
	return &MaxHeap[T]{
		BaseHeap: BaseHeap[T]{
			Data:       []T{},
			Type:       MaxHeapType,
			getScoreFn: getScoreFn,
		},
	}
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
