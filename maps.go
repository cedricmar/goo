package goo

import "sort"

// EachMap loops through a map, order is not guaranteed
func EachMap[U Ordered, T any](m map[U]T, fn func(el T, k U, m map[U]T)) {
	for k, el := range m {
		fn(el, k, m)
	}
}

// EachMapOrder loops through a map in the keys alphanumeric order
func EachMapOrdered[U Ordered, T any](m map[U]T, fn func(el T, k U, m map[U]T)) {
	ks := KeysOrdered(m)
	for k := range ks {
		fn(m[ks[k]], ks[k], m)
	}
}

// Keys collects the map keys
func Keys[U Scalar, T any](m map[U]T) (ks []U) {
	for k := range m {
		ks = append(ks, k)
	}
	return
}

// Keys collects the map keys and sorts them
func KeysOrdered[U Ordered, T any](m map[U]T) []U {
	ks := Keys(m)
	sort.Slice(ks, func(i, j int) bool {
		return ks[i] < ks[j]
	})
	return ks
}

// ReduceMap makes a map into a single value
func ReduceMap[U Ordered, T any, V any](m map[U]T, fn func(mem V, k U, el T) V) V {
	var mem V
	for k, el := range m {
		mem = fn(mem, k, el)
	}
	return mem
}

// FilterMap returns a new map with the "true" elements from the filter function
func FilterMap[U Ordered, T any](m map[U]T, fn func(k U, el T) bool) map[U]T {
	r := map[U]T{}
	for k, el := range m {
		if fn(k, el) {
			r[k] = el
		}
	}
	return r
}

// CloneMap allocates a copy of the given map
func CloneMap[U Ordered, T any](m map[U]T) map[U]T {
	r := map[U]T{}
	for k, el := range m {
		r[k] = el
	}
	return r
}
