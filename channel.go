package go2linq

import (
	"iter"
)

// ChanAll returns an [iterator] over the elements of the [channel].
// If 'c' is nil, [iterator] over the empty [sequence] is returned.
//
// [channel]: https://go.dev/ref/spec#Channel_types
// [iterator]: https://pkg.go.dev/iter#hdr-Iterators
// [sequence]: https://pkg.go.dev/iter#Seq
func ChanAll[E any](c <-chan E) iter.Seq[E] {
	if c == nil {
		return Empty[E]()
	}
	return func(yield func(E) bool) {
		for e := range c {
			if !yield(e) {
				return
			}
		}
	}
}

// ChanAll2 returns an [iterator] over index-element pairs of the [channel].
// If 'c' is nil, [iterator] over the empty [sequence] of pairs is returned.
//
// [channel]: https://go.dev/ref/spec#Channel_types
// [iterator]: https://pkg.go.dev/iter#hdr-Iterators
// [sequence]: https://pkg.go.dev/iter#Seq2
func ChanAll2[E any](c <-chan E) iter.Seq2[int, E] {
	if c == nil {
		return Empty2[int, E]()
	}
	return func(yield func(int, E) bool) {
		i := 0
		for e := range c {
			if !yield(i, e) {
				return
			}
			i++
		}
	}
}
