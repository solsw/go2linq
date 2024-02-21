package go2linq

import (
	"iter"
)

// [Reverse] inverts the order of the elements in a sequence.
//
// [Reverse]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.reverse
func Reverse[Source any](source iter.Seq[Source]) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return func(yield func(Source) bool) {
			ss, _ := ToSlice(source)
			for i := len(ss) - 1; i >= 0; i-- {
				if !yield(ss[i]) {
					return
				}
			}
		},
		nil
}
