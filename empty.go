package go2linq

import (
	"iter"
)

// [Empty] returns an empty sequence that has a specified type argument.
//
// [Empty]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.empty
func Empty[Result any]() iter.Seq[Result] {
	return func(func(Result) bool) {}
}

// [Empty2] returns an empty sequence2 that has a specified type arguments.
func Empty2[K, V any]() iter.Seq2[K, V] {
	return func(func(K, V) bool) {}
}
