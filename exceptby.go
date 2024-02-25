package go2linq

import (
	"iter"
	"sort"
	"sync"

	"github.com/solsw/generichelper"
)

// [ExceptBy] produces the set difference of two sequences according to
// a specified key selector function and using [generichelper.DeepEqual] as key equaler.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby
func ExceptBy[Source, Key any](first iter.Seq[Source], second iter.Seq[Key], keySelector func(Source) Key) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ExceptByEq(first, second, keySelector, generichelper.DeepEqual[Key])
}

// [ExceptByEq] produces the set difference of two sequences according to
// a specified key selector function and using a specified key equaler.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby
func ExceptByEq[Source, Key any](first iter.Seq[Source], second iter.Seq[Key],
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
	return func(yield func(Source) bool) {
			distinct1, _ := Distinct(first)
			var once sync.Once
			var distinct2 []Key
			for s := range distinct1 {
				once.Do(func() { deq2, _ := DistinctEq(second, equal); distinct2, _ = ToSlice(deq2) })
				k := keySelector(s)
				if !elInElelEq(k, distinct2, equal) {
					if !yield(s) {
						return
					}
				}
			}
		},
		nil
}

// [ExceptByCmp] produces the set difference of two sequences according to a specified
// key selector function and using a specified 'compare' to compare keys. (See [DistinctCmp].)
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby
func ExceptByCmp[Source, Key any](first iter.Seq[Source], second iter.Seq[Key],
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
	return func(yield func(Source) bool) {
			distinct1, _ := Distinct(first)
			var once2 sync.Once
			var distinct2 []Key
			for s := range distinct1 {
				once2.Do(func() {
					deq2, _ := DistinctCmp(second, compare)
					distinct2, _ = ToSlice(deq2)
					sort.Slice(distinct2, func(i, j int) bool { return compare(distinct2[i], distinct2[j]) < 0 })
				})
				k := keySelector(s)
				if !elInElelCmp(k, distinct2, compare) {
					if !yield(s) {
						return
					}
				}
			}
		},
		nil
}
