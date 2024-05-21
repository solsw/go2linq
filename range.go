package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Range] generates a sequence of [int]s within a specified range.
//
// [int]: https://pkg.go.dev/builtin#int
// [Range]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.range
func Range(start, count int) (iter.Seq[int], error) {
	if count < 0 {
		return nil, errorhelper.CallerError(ErrNegativeCount)
	}
	return func(yield func(int) bool) {
			for i := range count {
				if !yield(start + i) {
					return
				}
			}
		},
		nil
}
