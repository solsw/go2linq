//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 14 - Distinct
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-14-distinct/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

// Distinct returns distinct elements from a sequence by using reflect.DeepEqual to compare values.
// Distinct panics if 'source' is nil.
func Distinct[Source any](source Enumerator[Source]) Enumerator[Source] {
	return DistinctEq(source, nil)
}

// DistinctErr is like Distinct but returns an error instead of panicking.
func DistinctErr[Source any](source Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Distinct(source), nil
}

// DistinctEq returns distinct elements from a sequence by using a specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// DistinctEq panics if 'source' is nil.
func DistinctEq[Source any](source Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	var c Source
	seen := make([]Source, 0)
	return OnFunc[Source]{
		MvNxt: func() bool {
			for source.MoveNext() {
				c = source.Current()
				if !elInElelEq(c, seen, eq) {
					seen = append(seen, c)
					return true
				}
			}
			return false
		},
		Crrnt: func() Source { return c },
		Rst:   func() { seen = make([]Source, 0); source.Reset() },
	}
}

// DistinctEqErr is like DistinctEq but returns an error instead of panicking.
func DistinctEqErr[Source any](source Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return DistinctEq(source, eq), nil
}

// DistinctCmp returns distinct elements from a sequence by using a specified Comparer to compare values.
// DistinctCmpErr panics if 'source' or 'cmp' is nil.
//
// Sorted slice of already seen elements is internally built.
// Sorted slice allows to use binary search to determine whether the element was seen or not.
// This may give performance gain when processing large sequences (though this is a subject for benchmarking).
func DistinctCmp[Source any](source Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if cmp == nil {
		panic(ErrNilComparer)
	}
	var c Source
	seen := make([]Source, 0)
	return OnFunc[Source]{
		MvNxt: func() bool {
			for source.MoveNext() {
				c = source.Current()
				i := elIdxInElelCmp(c, seen, cmp)
				if i == len(seen) || cmp.Compare(c, seen[i]) != 0 {
					elIntoElelAtIdx(c, &seen, i)
					return true
				}
			}
			return false
		},
		Crrnt: func() Source { return c },
		Rst:   func() { seen = make([]Source, 0); source.Reset() },
	}
}

// DistinctCmpErr is like DistinctCmp but returns an error instead of panicking.
func DistinctCmpErr[Source any](source Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return DistinctCmp(source, cmp), nil
}
