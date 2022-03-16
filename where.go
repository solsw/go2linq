//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 2 - "Where"
// https://codeblog.jonskeet.uk/2010/09/03/reimplementing-linq-to-objects-part-2-quot-where-quot/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func enrWhere[Source any](source Enumerable[Source], predicate func(Source) bool) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = enr.Current()
					if predicate(c) {
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enr.Reset() },
		}
	}
}

// Where filters a sequence of values based on a predicate.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where)
func Where[Source any](source Enumerable[Source], predicate func(Source) bool) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return OnFactory(enrWhere(source, predicate)), nil
}

// WhereMust is like Where but panics in case of error.
func WhereMust[Source any](source Enumerable[Source], predicate func(Source) bool) Enumerable[Source] {
	r, err := Where(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

func enrWhereIdx[Source any](source Enumerable[Source], predicate func(Source, int) bool) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enr := source.GetEnumerator()
		var c Source
		i := -1 // position before the first element
		return enrFunc[Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = enr.Current()
					i++
					if predicate(c, i) {
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { i = -1; enr.Reset() },
		}
	}
}

// WhereIdx filters a sequence of values based on a predicate.
// Each element's index is used in the logic of the predicate function.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where)
func WhereIdx[Source any](source Enumerable[Source], predicate func(Source, int) bool) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return OnFactory(enrWhereIdx(source, predicate)), nil
}

// WhereIdxMust is like WhereIdx but panics in case of error.
func WhereIdxMust[Source any](source Enumerable[Source], predicate func(Source, int) bool) Enumerable[Source] {
	r, err := WhereIdx(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
