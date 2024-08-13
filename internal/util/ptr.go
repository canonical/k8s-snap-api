package util

// Deref returns the concrete value of a pointer, or the zero value for that type.
func Deref[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var zero T
	return zero
}

// PtrTo returns a pointer to a concrete value.
func PtrTo[T any](val T) *T {
	return &val
}
