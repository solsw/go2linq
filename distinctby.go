package go2linq

import (
	"github.com/solsw/collate"
)

// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby

// [DistinctBy] returns distinct elements from a sequence according to a specified key selector function
// and using [collate.DeepEqualer] to compare keys.
//
// [DistinctBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
func DistinctBy[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return DistinctByEq(source, keySelector, nil)
}

// DistinctByMust is like [DistinctBy] but panics in case of error.
func DistinctByMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) Enumerable[Source] {
	r, err := DistinctBy(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

func factoryDistinctByEq[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler collate.Equaler[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		var c Source
		seen := make([]Key, 0)
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = enr.Current()
					k := keySelector(c)
					if !elInElelEq(k, seen, equaler) {
						seen = append(seen, k)
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { seen = make([]Key, 0); enr.Reset() },
		}
	}
}

// [DistinctByEq] returns distinct elements from a sequence according to a specified key selector function
// and using a specified equaler to compare keys. If 'equaler' is nil [collate.DeepEqualer] is used.
//
// [DistinctByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
func DistinctByEq[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler collate.Equaler[Key]) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	return OnFactory(factoryDistinctByEq(source, keySelector, equaler)), nil
}

// DistinctByEqMust is like [DistinctByEq] but panics in case of error.
func DistinctByEqMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, equaler collate.Equaler[Key]) Enumerable[Source] {
	r, err := DistinctByEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

func factoryDistinctByCmp[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, comparer collate.Comparer[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		var c Source
		seen := make([]Key, 0)
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = enr.Current()
					k := keySelector(c)
					i := elIdxInElelCmp(k, seen, comparer)
					if i == len(seen) || comparer.Compare(k, seen[i]) != 0 {
						elIntoElelAtIdx(k, &seen, i)
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { seen = make([]Key, 0); enr.Reset() },
		}
	}
}

// [DistinctByCmp] returns distinct elements from a sequence according to a specified key selector function
// and using a specified comparer to compare keys. (See [DistinctCmp].)
//
// [DistinctByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
func DistinctByCmp[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, comparer collate.Comparer[Key]) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return OnFactory(factoryDistinctByCmp(source, keySelector, comparer)), nil
}

// DistinctByCmpMust is like [DistinctByCmp] but panics in case of error.
func DistinctByCmpMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key, comparer collate.Comparer[Key]) Enumerable[Source] {
	r, err := DistinctByCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
