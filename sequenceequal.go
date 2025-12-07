package go2linq

import (
	"iter"

	"github.com/solsw/iterhelper"
)

// [SequenceEqual] determines whether two [sequences] are equal
// by comparing their elements using [generichelper.DeepEqual].
//
// [SequenceEqual]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
// [sequences]: https://pkg.go.dev/iter#Seq
func SequenceEqual[Source any](first, second iter.Seq[Source]) (bool, error) {
	return iterhelper.SeqEqual(first, second)
}

// [SequenceEqualEq] determines whether two [sequences] are equal
// by comparing their elements using a specified function.
//
// [SequenceEqualEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
// [sequences]: https://pkg.go.dev/iter#Seq
func SequenceEqualEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (bool, error) {
	return iterhelper.SeqEqualEq(first, second, equal)
}
