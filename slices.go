package goo

// Each loops through a slice
func Each[T any](s []T, fn func(el T, i int, s []T)) {
	for i, el := range s {
		fn(el, i, s)
	}
}

// Map each values through a modifier
func Map[T any](s []T, fn func(el T) T) {
	for i, el := range s {
		s[i] = fn(el)
	}
}

// Reduce makes a slice into a single value
func Reduce[T, U any](s []T, fn func(mem U, el T) U) U {
	var mem U
	for _, el := range s {
		mem = fn(mem, el)
	}
	return mem
}

// Clone a slice
func Clone[T any](s []T) []T {
	c := make([]T, len(s), cap(s))
	copy(c, s)
	return c
}

// Unique removes identical elements from a slice
func Unique[T Scalar](s []T) []T {
	u := LookupTable(s)
	return Keys(u)
}

// Find the first "true" element from the test function
func Find[T any](s []T, fn func(el T) bool) (r T) {
	for _, el := range s {
		if fn(el) {
			return el
		}
	}
	return r
}

// Filter returns a new slice with the "true" elements from the filter function
func Filter[T any](s []T, fn func(el T) bool) []T {
	var r []T
	for _, el := range s {
		if fn(el) {
			r = append(r, el)
		}
	}
	return r
}

// FoundIn returns a boolean if the element k was found in the slice s
func FoundIn[T comparable](s []T, k T) bool {
	_, f := FoundAt(s, k)
	return f
}

// FoundAt returns the index of the found element in a slice s or -1
// It also returns a boolean to indicate if the element k was found
func FoundAt[T comparable](s []T, k T) (int, bool) {
	for i, el := range s {
		if k == el {
			return i, true
		}
	}
	return -1, false
}

// Index turns a slice into a map using a function on each slice element to infer its key
func Index[K comparable, T any](s []T, fn func(el T) K) map[K]T {
	m := make(map[K]T)
	for _, el := range s {
		m[fn(el)] = el
	}
	return m
}

// LookupTable inserts slice elements as keys in the returned map
func LookupTable[T Scalar](s []T) map[T]struct{} {
	lt := map[T]struct{}{}
	Each(s, func(el T, i int, s []T) {
		lt[el] = struct{}{}
	})
	return lt
}
