package maps

// GetOrDefault is a utility function
// similar to java's getOrDefault
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

// MapOf returns a map from the given key-value pairs.
// The input must be even-length: alternating keys and values.
// If the number of arguments is odd, it panics.
//
// Example:
//
//	m := MapOf("k1", 1, "k2", 2)  // map[string]int{"k1": 1, "k2": 2}
func MapOf[K comparable, V any](pairs ...any) map[K]V {
	if len(pairs)%2 != 0 {
		panic("MapOf requires even number of arguments: alternating keys and values")
	}

	m := make(map[K]V, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		k, ok1 := pairs[i].(K)
		v, ok2 := pairs[i+1].(V)
		if !ok1 || !ok2 {
			panic("MapOf: type assertion failed")
		}
		m[k] = v
	}
	return m
}

// SetOf returns a set represented as a map[T]struct{} containing the provided elements.
// Duplicate values are automatically deduplicated.
//
// Example:
//
//	s := SetOf("a", "b", "a")  // map[string]struct{}{"a": {}, "b": {}}
func SetOf[T comparable](elements ...T) map[T]struct{} {
	set := make(map[T]struct{}, len(elements))
	for _, e := range elements {
		set[e] = struct{}{}
	}
	return set
}
