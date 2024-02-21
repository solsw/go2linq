package go2linq

import (
	"cmp"
	"iter"
	"sort"
)

func orderByKeyLsPrim[Source, Key any](source iter.Seq[Source],
	keySelector func(Source) Key, less func(x, y Key) bool) iter.Seq[Source] {
	type sk struct {
		s Source
		k Key
	}
	var sksk []sk
	for s := range source {
		sksk = append(sksk, sk{s: s, k: keySelector(s)})
	}
	sort.SliceStable(sksk, func(i, j int) bool {
		return less(sksk[i].k, sksk[j].k)
	})
	ss := make([]Source, len(sksk))
	for i := range len(sksk) {
		ss[i] = sksk[i].s
	}
	return SliceAll(ss)
}

// [OrderByKey] sorts the elements of a sequence in ascending order according to a key.
//
// [OrderByKey]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func OrderByKey[Source any, Key cmp.Ordered](source iter.Seq[Source],
	keySelector func(Source) Key) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return orderByKeyLsPrim(source, keySelector, cmp.Less), nil
}

// [OrderByKeyLs] sorts the elements of a sequence in ascending order of keys using a specified 'less' function.
//
// [OrderByKeyLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderby
func OrderByKeyLs[Source, Key any](source iter.Seq[Source],
	keySelector func(Source) Key, less func(x, y Key) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if less == nil {
		return nil, ErrNilLess
	}
	return orderByKeyLsPrim(source, keySelector, less), nil
}

// [OrderByKeyDesc] sorts the elements of a sequence in descending order according to a key.
//
// [OrderByKeyDesc]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func OrderByKeyDesc[Source any, Key cmp.Ordered](source iter.Seq[Source],
	keySelector func(Source) Key) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return orderByKeyLsPrim(source, keySelector, reverseLess[Key](cmp.Less)), nil
}

// [OrderByKeyDescLs] sorts the elements of a sequence in descending order of keys using a 'less' function.
//
// [OrderByKeyDescLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
func OrderByKeyDescLs[Source, Key any](source iter.Seq[Source],
	keySelector func(Source) Key, less func(x, y Key) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if less == nil {
		return nil, ErrNilLess
	}
	return orderByKeyLsPrim(source, keySelector, reverseLess[Key](less)), nil
}
