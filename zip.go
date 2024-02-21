package go2linq

import (
	"iter"
)

// [Zip] applies a specified function to the corresponding elements
// of two sequences, producing a sequence of the results.
//
// [Zip]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.zip
func Zip[First, Second, Result any](first iter.Seq[First], second iter.Seq[Second],
	resultSelector func(First, Second) Result) (iter.Seq[Result], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if resultSelector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			next1, stop1 := iter.Pull(first)
			defer stop1()
			next2, stop2 := iter.Pull(second)
			defer stop2()
			for {
				f, ok1 := next1()
				s, ok2 := next2()
				if !ok1 || !ok2 {
					return
				}
				if !yield(resultSelector(f, s)) {
					return
				}
			}
		},
		nil
}
