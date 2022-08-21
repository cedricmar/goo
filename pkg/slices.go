package goo

// Each loops through a slice
func Each[T any](s []T, fn func(el T, i int, s []T)) {
	for i, el := range s {
		fn(el, i, s)
	}
}
