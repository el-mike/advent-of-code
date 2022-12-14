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

// AnySatisfies - returns true if for any value in the map callback returns true,
// false otherwise.
func AnySatisfies[TKey comparable, TValue comparable](
	m map[TKey]TValue,
	cb func(key TKey, value TValue) bool,
) bool {
	for key, value := range m {
		if cb(key, value) {
			return true
		}
	}

	return false
}

func CountWhere[TKey comparable, TValue comparable](
	m map[TKey]TValue,
	cb func(key TKey, value TValue) bool,
) int {
	count := 0
	for key, value := range m {
		if cb(key, value) {
			count += 1
		}
	}

	return count
}

// CopyMap - returns a shallow copy of a given map.
func CopyMap[TKey comparable, TValue any](m map[TKey]TValue) map[TKey]TValue {
	mapCopy := map[TKey]TValue{}

	for key, value := range m {
		mapCopy[key] = value
	}

	return mapCopy
}
