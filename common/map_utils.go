package common

// GetKeyByValue - returns a key from the map based on given value.
// If does not exist, returns nil pointer for the generic type.
func GetKeyByValue[TKey comparable, TValue comparable](m map[TKey]TValue, target TValue) TKey {
	for key, value := range m {
		if value == target {
			return key
		}
	}

	var nilResult TKey

	return nilResult
}
