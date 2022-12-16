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
