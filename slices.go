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

func Index[K comparable, T any](s []T, fn func(el T) K) map[K]T {
	m := make(map[K]T)
	for _, el := range s {
		m[fn(el)] = el
	}
	return m
}

// Unique removes identical elements from a slice
func Unique[T Scalar](s []T) []T {
	u := LookupTable(s)
	return Keys(u)
}

// LookupTable inserts slice elements as keys in a map
func LookupTable[T Scalar](s []T) map[T]struct{} {
	lt := map[T]struct{}{}
	Each(s, func(el T, i int, s []T) {
		lt[el] = struct{}{}
	})
	return lt
}

// FoundAt returns the index of the found element in a slice s or -1
// It also returns a boolean to indicate if the element k was found
func FoundAt[T comparable](k T, s []T) (int, bool) {
	for i, el := range s {
		if k == el {
			return i, true
		}
	}
	return -1, false
}

// FoundIn returns a boolean if the element k was found in the slice s
func FoundIn[T comparable](k T, s []T) bool {
	_, f := FoundAt(k, s)
	return f
}
