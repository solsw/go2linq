package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [UnionBy] produces the set union of two sequences according to
// a specified key selector function and using [generichelper.DeepEqual] as key equaler.
//
// [UnionBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionBy[Source, Key any](first, second iter.Seq[Source], keySelector func(Source) Key) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return UnionByEq(first, second, keySelector, generichelper.DeepEqual[Key])
}

// [UnionByEq] produces the set union of two sequences according to
// a specified key selector function and using a specified key equaler.
//
// [UnionByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionByEq[Source, Key any](first, second iter.Seq[Source],
	keySelector func(Source) Key, equal func(Key, Key) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	concat, _ := Concat(first, second)
	return DistinctByEq(concat, keySelector, equal)
}

// [UnionByCmp] produces the set union of two sequences according to a specified
// key selector function and using a specified key comparer. (See [DistinctCmp].)
//
// [UnionByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionByCmp[Source, Key any](first, second iter.Seq[Source],
	keySelector func(Source) Key, compare func(Key, Key) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if compare == nil {
		return nil, ErrNilCompare
	}
	concat, _ := Concat(first, second)
	return DistinctByCmp(concat, keySelector, compare)
}
