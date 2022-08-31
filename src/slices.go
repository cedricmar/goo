package goo

// Each loops through a slice
func Each[T any](s []T, fn func(el T, i int, s []T)) {
	for i, el := range s {
		fn(el, i, s)
	}
}

// Map each values through a modifier (func)
func Map[T any](s []T, fn func(el T) T) {
	for i, el := range s {
		s[i] = fn(el)
	}
}

// Clone a slice
func Clone[T any](s []T) []T {
	c := make([]T, len(s), cap(s))
	copy(c, s)
	return c
}
