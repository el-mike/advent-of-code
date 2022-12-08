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
