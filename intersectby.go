package go2linq

import (
	"iter"
	"sort"
	"sync"

	"github.com/solsw/generichelper"
)

// [IntersectBy] produces the set intersection of two sequences according to
// a specified key selector function and using [generichelper.DeepEqual] as key equaler.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [IntersectBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersectby
func IntersectBy[Source, Key any](first iter.Seq[Source], second iter.Seq[Key], keySelector func(Source) Key) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return IntersectByEq(first, second, keySelector, generichelper.DeepEqual[Key])
}

func seqIntersectByEq[Source, Key any](first iter.Seq[Source], second iter.Seq[Key],
	keySelector func(Source) Key, sourceEqual func(Source, Source) bool,
	keyEqual func(Key, Key) bool, keyCompare func(Key, Key) int) func(func(Source) bool) {
	return func(yield func(Source) bool) {
		d1, _ := DistinctEq(first, sourceEqual)
		var once sync.Once
		var sl2 []Key
		for s := range d1 {
			once.Do(func() {
				if keyCompare == nil {
					d2, _ := DistinctEq(second, keyEqual)
					sl2, _ = ToSlice(d2)
				} else {
					d2, _ := DistinctCmp(second, keyCompare)
					sl2, _ = ToSlice(d2)
					sort.Slice(sl2, func(i, j int) bool { return keyCompare(sl2[i], sl2[j]) < 0 })
				}
			})
			k := keySelector(s)
			if keyCompare == nil {
				if elInElelEq(k, sl2, keyEqual) {
					if !yield(s) {
						return
					}
				}
			} else {
				if elInElelCmp(k, sl2, keyCompare) {
					if !yield(s) {
						return
					}
				}
			}
		}
	}
}

// [IntersectByEq] produces the set intersection of two sequences according to
// a specified key selector function and using a specified key equaler.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [IntersectByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersectby
func IntersectByEq[Source, Key any](first iter.Seq[Source], second iter.Seq[Key],
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
	return seqIntersectByEq(first, second, keySelector, generichelper.DeepEqual[Source], equal, nil),
		nil
}

// [IntersectByCmp] produces the set intersection of two sequences according to
// a specified key selector function and using a specified key comparer. (See [DistinctCmp].)
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [IntersectByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersectby
func IntersectByCmp[Source, Key any](first iter.Seq[Source], second iter.Seq[Key],
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
	return seqIntersectByEq(first, second, keySelector, generichelper.DeepEqual[Source], nil, compare),
		nil
}
