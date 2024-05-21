package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Repeat] generates a sequence that contains one repeated value.
//
// [Repeat]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.repeat
func Repeat[Result any](element Result, count int) (iter.Seq[Result], error) {
	if count < 0 {
		return nil, errorhelper.CallerError(ErrNegativeCount)
	}
	return func(yield func(Result) bool) {
			for range count {
				if !yield(element) {
					return
				}
			}
		},
		nil
}
