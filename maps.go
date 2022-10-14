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
