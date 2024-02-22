package go2linq

import (
	"iter"
)

// MapAll returns an iterator over all entries of the [map].
// If 'm' is nil, empty iterator is returned.
//
// [map]: https://go.dev/ref/spec#Map_types
func MapAll[K comparable, E any](m map[K]E) iter.Seq2[K, E] {
	return func(yield func(K, E) bool) {
		for k, e := range m {
			if !yield(k, e) {
				return
			}
		}
	}
}
