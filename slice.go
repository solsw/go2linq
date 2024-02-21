package go2linq

import (
	"iter"
)

// SliceAll returns an iterator over values of all elements of the [slice].
// If 's' is nil, empty iterator is returned.
//
// [slice]: https://go.dev/ref/spec#Slice_types
func SliceAll[S ~[]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range s {
			if !yield(e) {
				return
			}
		}
	}
}

// VarAll returns an iterator over all elements of the [variadic] [slice].
//
// [variadic]: https://go.dev/ref/spec#Function_types
// [slice]: https://go.dev/ref/spec#Slice_types
func VarAll[E any](s ...E) iter.Seq[E] {
	return SliceAll(s)
}

// SliceAll2 returns an iterator over pairs of index and value of all elements of the [slice].
// If 's' is nil, empty iterator is returned.
//
// [slice]: https://go.dev/ref/spec#Slice_types
func SliceAll2[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, e := range s {
			if !yield(i, e) {
				return
			}
		}
	}
}
