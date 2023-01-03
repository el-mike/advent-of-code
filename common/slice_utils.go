package common

// Contains - returns true if slice contains given value, false otherwise.
func Contains[T comparable](slice []T, value T) bool {
	for _, x := range slice {
		if x == value {
			return true
		}
	}

	return false
}

func RemoveDuplicates[T comparable](slice []T) []T {
	var result []T

	for _, x := range slice {
		if !Contains[T](result, x) {
			result = append(result, x)
		}
	}

	return result
}

func Filter[T comparable](slice []T, predictor func(x T) bool) []T {
	var result []T

	for _, x := range slice {
		if predictor(x) {
			result = append(result, x)
		}
	}

	return result
}

func RemoveAt[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func InsertAt[T comparable](slice []T, value T, index int) []T {
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

func Move[T comparable](slice []T, srcIndex, dstIndex int) []T {
	// We need to extract the value before we call RemoveAt.
	value := slice[srcIndex]

	return InsertAt[T](RemoveAt(slice, srcIndex), value, dstIndex)
}
