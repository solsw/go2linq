//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 14 - Distinct
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-14-distinct/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

// Distinct returns distinct elements from a sequence by using reflect.DeepEqual to compare values.
func Distinct[Source any](source Enumerator[Source]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return DistinctEq(source, nil)
}

// DistinctMust is like Distinct but panics in case of error.
func DistinctMust[Source any](source Enumerator[Source]) Enumerator[Source] {
	r, err := Distinct(source)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctEq returns distinct elements from a sequence by using a specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
func DistinctEq[Source any](source Enumerator[Source], eq Equaler[Source]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	var c Source
	seen := make([]Source, 0)
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
					if !elInElelEq(c, seen, eq) {
						seen = append(seen, c)
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { seen = make([]Source, 0); source.Reset() },
		},
		nil
}

// DistinctEqMust is like DistinctEq but panics in case of error.
func DistinctEqMust[Source any](source Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	r, err := DistinctEq(source, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctCmp returns distinct elements from a sequence by using a specified Comparer to compare values.
//
// Sorted slice of already seen elements is internally built.
// Sorted slice allows to use binary search to determine whether the element was seen or not.
// This may give performance gain when processing large sequences (though this is a subject for benchmarking).
func DistinctCmp[Source any](source Enumerator[Source], comparer Comparer[Source]) (Enumerator[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	var c Source
	seen := make([]Source, 0)
	return OnFunc[Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = source.Current()
					i := elIdxInElelCmp(c, seen, comparer)
					if i == len(seen) || comparer.Compare(c, seen[i]) != 0 {
						elIntoElelAtIdx(c, &seen, i)
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { seen = make([]Source, 0); source.Reset() },
		},
		nil
}

// DistinctCmpMust is like DistinctCmp but panics in case of error.
func DistinctCmpMust[Source any](source Enumerator[Source], comparer Comparer[Source]) Enumerator[Source] {
	r, err := DistinctCmp(source, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
