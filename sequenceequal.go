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
	return iterhelper.SequenceEqual(first, second)
}

// [SequenceEqualEq] determines whether two [sequences] are equal
// by comparing their elements using a specified function.
//
// [SequenceEqualEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
// [sequences]: https://pkg.go.dev/iter#Seq
func SequenceEqualEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (bool, error) {
	return iterhelper.SequenceEqualEq(first, second, equal)
}

// SequenceEqual2 determines whether two [sequences] are equal
// by comparing their elements using [generichelper.DeepEqual].
//
// [sequences]: https://pkg.go.dev/iter#Seq2
func SequenceEqual2[K, V any](first, second iter.Seq2[K, V]) (bool, error) {
	return iterhelper.SequenceEqual2(first, second)
}

// SequenceEqual2Eq determines whether two [sequences] are equal
// by comparing their elements using specified equals.
//
// [sequences]: https://pkg.go.dev/iter#Seq2
func SequenceEqual2Eq[K, V any](first, second iter.Seq2[K, V], equalK func(K, K) bool, equalV func(V, V) bool) (bool, error) {
	return iterhelper.SequenceEqual2Eq(first, second, equalK, equalV)
}
