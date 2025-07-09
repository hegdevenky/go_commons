package arrays

// Of returns a slice containing the provided elements in order.
// This is similar to Java's List.of() but does not guarantee immutability.
//
// Note: The returned slice may be modified by the caller.
// For read-only behavior, the caller should avoid modifying it.
//
// Example:
//
//	s := Of(1, 2, 3)  // []int{1, 2, 3}
func Of[T any](elems ...T) []T {
	return elems
}

// CopyOf returns a deep copy of the provided slice or vararg elements.
// The resulting slice has the same contents as the input but is a new allocation.
// It ensures that modifications to the result do not affect the original.
//
// Example:
//
//	original := []string{"a", "b", "c"}
//	copy := CopyOf(original...)  // modifies 'copy' without affecting 'original'
func CopyOf[T any](elems ...T) []T {
	cp := make([]T, len(elems))
	copy(cp, elems)
	return cp
}

// Fill assigns the specified value to each element of the specified slice.
//
// Example:
// slice1 := make([]int, 4) // slice1 = [0,0,0,0]
// slice1 = Fill(slice1, 123) // slice1 = [123,123,123,123]
//
// This function modifies the given slice, so need not reassign the returned slice.
func Fill[T any](a []T, value T) []T {
	for i, _ := range a {
		a[i] = value
	}
	return a
}
