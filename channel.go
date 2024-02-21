package go2linq

import (
	"iter"
)

// ChanToSeq converts a [channel] to a sequence.
// If 'c' is nil, empty sequence is returned.
//
// [channel]: https://go.dev/ref/spec#Channel_types
func ChanToSeq[E any](c <-chan E) iter.Seq[E] {
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

// ChanToSeq2 converts a [channel] to a sequence2.
// If 'c' is nil, empty sequence2 is returned.
//
// [channel]: https://go.dev/ref/spec#Channel_types
func ChanToSeq2[E any](c <-chan E) iter.Seq2[int, E] {
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
