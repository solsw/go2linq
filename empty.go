package go2linq

import (
	"iter"

	"github.com/solsw/iterhelper"
)

// [Empty] returns an [iterator] over an empty [sequence] of values.
//
// [Empty]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.empty
// [iterator]: https://pkg.go.dev/iter#hdr-Iterators
// [sequence]: https://pkg.go.dev/iter#Seq
func Empty[Result any]() iter.Seq[Result] {
	return iterhelper.Empty[Result]()
}
