package gotask

func zero[T any]() T {
	var z T

	return z
}

func isZero[T comparable](value T) bool {
	return value == zero[T]()
}
