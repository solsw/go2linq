//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinctby

// DistinctBy returns distinct elements from a sequence according to a specified key selector function
// and using reflect.DeepEqual to compare keys.
func DistinctBy[Source, Key any](source Enumerator[Source], keySelector func(Source) Key) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return DistinctByEq(source, keySelector, nil)
}

// DistinctByMust is like DistinctBy but panics in case of error.
func DistinctByMust[Source, Key any](source Enumerator[Source], keySelector func(Source) Key) Enumerator[Source] {
	r, err := DistinctBy(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctByEq returns distinct elements from a sequence according to a specified key selector function
// and using a specified equaler to compare keys. If 'equaler' is nil reflect.DeepEqual is used.
func DistinctByEq[Source, Key any](source Enumerator[Source], keySelector func(Source) Key, equaler Equaler[Key]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = EqualerFunc[Key](DeepEqual[Key])
	}
	var c Source
	seen := make([]Key, 0)
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
					k := keySelector(c)
					if !elInElelEq(k, seen, equaler) {
						seen = append(seen, k)
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { seen = make([]Key, 0); source.Reset() },
		},
		nil
}

// DistinctByEqMust is like DistinctByEq but panics in case of error.
func DistinctByEqMust[Source, Key any](source Enumerator[Source], keySelector func(Source) Key, equaler Equaler[Key]) Enumerator[Source] {
	r, err := DistinctByEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctByCmp returns distinct elements from a sequence according to a specified key selector function
// and using a specified comparer to compare keys.
func DistinctByCmp[Source, Key any](source Enumerator[Source], keySelector func(Source) Key, comparer Comparer[Key]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	var c Source
	seen := make([]Key, 0)
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
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
			rst:   func() { seen = make([]Key, 0); source.Reset() },
		},
		nil
}

// DistinctByCmpMust is like DistinctByCmp but panics in case of error.
func DistinctByCmpMust[Source, Key any](source Enumerator[Source], keySelector func(Source) Key, comparer Comparer[Key]) Enumerator[Source] {
	r, err := DistinctByCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
