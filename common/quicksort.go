package common

func QuickSort(slice []int, l int, r int) []int {
	if l >= r {
		return slice
	}

	pivot := partition(slice, l, r)

	QuickSort(slice, l, pivot-1)
	QuickSort(slice, pivot+1, r)

	return slice
}

func partition(slice []int, l int, r int) int {
	pivot := slice[r]
	i := l - 1

	for j := l; j < r; j += 1 {
		if slice[j] <= pivot {
			i += 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[r] = slice[r], slice[i+1]

	return i + 1
}
